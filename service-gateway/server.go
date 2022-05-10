package main

// The server uses environment variables. If they are not set the server won't start. It expects the following environment variables:
//   * SERVER_PORT_SSL       - the server is listening on this portnumber and starts an HTTPS-Server
//   * ROUTES_CONFIG         - path and filename where to find the routes.json config file to define all known routes

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"

	"github.com/damonkeys/covid-checkin/monkeys/config"
	l "github.com/damonkeys/covid-checkin/monkeys/logger"
	mcmiddleware "github.com/damonkeys/covid-checkin/monkeys/middleware"
	"github.com/damonkeys/covid-checkin/monkeys/tracing"
)

// RoutesStruct defines all configured routes
type RoutesStruct struct {
	Routes []RouteStruct `json:"routes"`
}

// RouteStruct defines one single route. The given path will be routed to one of the given URLs.
type RouteStruct struct {
	Name         string   `json:"name"`
	Path         string   `json:"path"`
	Description  string   `json:"description"`
	Urls         []string `json:"urls"`
	Auth         BasicAuthStruct
	LoadBalancer string `json:"balancer"`
	Rewrite      bool   `json:"rewrite"`
}

// ProxyConfigStruct defines the whole configuration for the reverse proxy. It will be filled via commandline parameters
type ProxyConfigStruct struct {
	ServerPort       string `env:"SERVER_PORT"`
	SSLActive        bool   `env:"SSL_ACTIVE"`
	RoutesConfigFile string `env:"ROUTES_CONFIG"`
}

// BasicAuthStruct contains information for basic auth: User and Password
type BasicAuthStruct struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// RouteAuthorizer creates a middleware function that authorizes against configured Username and Password
type RouteAuthorizer interface {
	authorize(ctx context.Context, e *echo.Echo, balancer middleware.ProxyBalancer)
}

func (route RouteStruct) simpleUserNamePasswordCheck(username, password string, c echo.Context) (bool, error) {
	span := tracing.Enter(c)
	defer span.Finish()
	if username == route.Auth.User && password == route.Auth.Password {
		tracing.LogString(span, "auth succssful", fmt.Sprintf("route: %s, user: %s", route.Name, route.Auth.User))
		return true, nil
	}
	tracing.LogString(span, "auth not succssful", fmt.Sprintf("route: %s, user: %s", route.Name, route.Auth.User))
	return false, nil
}

func (route RouteStruct) middlewareBasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(route.simpleUserNamePasswordCheck)
}

func (route RouteStruct) authorize(ctx context.Context, e *echo.Echo, balancer middleware.ProxyBalancer) {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()
	tracing.LogString(span, "route auth init", route.Name)
	e.Group(route.Path, route.middlewareBasicAuth(), middleware.Proxy(balancer))
}

// ProxyConfig defines the configuration for the reverse proxy
var proxyConfig ProxyConfigStruct

func main() {
	// tracer init
	closer, span, ctx := tracing.InitJaeger("service-gateway")
	defer closer.Close()

	setProxyConfig(ctx)

	// init echo
	e := echo.New()
	e.Use(tracing.MiddlewareWithoutCurrentUser("service-gateway"))
	if proxyConfig.SSLActive {
		e.Pre(middleware.HTTPSRedirect()) // Add https-redirect

		// needed for AutoTLS
		e.Use(middleware.Recover())
		e.AutoTLSManager.Cache = autocert.DirCache("./.cache")
	}

	// Logger config
	l.ConfigureLogger(ctx, "service-gateway", e)

	parseJSONConfig(ctx, e)

	// Start Proxy
	span.Finish()
	if proxyConfig.SSLActive {
		e.Logger.Fatal(e.StartAutoTLS(":" + proxyConfig.ServerPort))
	} else {
		e.Logger.Fatal(e.Start(":" + proxyConfig.ServerPort))
	}
}

// setProxyConfig gets environment variables and stores them into the ProxyConfigStruct
func setProxyConfig(ctx context.Context) {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	// read config from environment variables to struct
	configInterface, err := config.ReadEnvVars(ctx, ProxyConfigStruct{})
	if err != nil {
		tracing.LogError(span, err)
		log.Println(err)
		os.Exit(-1)
	}
	proxyConfig = configInterface.(ProxyConfigStruct)
}

// parseJSONConfig parses the json-routes config file and builds the routes for echo.
func parseJSONConfig(ctx context.Context, e *echo.Echo) {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	// parse config file
	var routes *RoutesStruct
	rewritePaths := make(map[string]string)

	data, err := ioutil.ReadFile(proxyConfig.RoutesConfigFile)
	if err != nil {
		e.Logger.Fatal("Cannot parse routes-config. Error: %s", err)
		tracing.LogError(span, err)
		os.Exit(0)
	}
	err = json.Unmarshal(data, &routes)
	tracing.LogStruct(span, "routesConfig", routes)

	// Setup proxy from struct
	for _, route := range routes.Routes {
		targets := []*middleware.ProxyTarget{}
		for _, serverURL := range route.Urls {
			urlParsed, err := url.Parse(serverURL)
			if err != nil {
				e.Logger.Fatal(err)
			}
			logString := fmt.Sprintf("Routing %s->%s", route.Description, urlParsed)
			e.Logger.Infof(logString)
			tracing.LogString(span, "Routing added", logString)
			targets = append(targets, &(middleware.ProxyTarget{Name: route.Name, URL: urlParsed}))
		}
		balancer := middleware.NewRoundRobinBalancer(targets)
		if strings.ToLower(route.LoadBalancer) == "random" {
			balancer = middleware.NewRandomBalancer(targets)
		}

		//routeGroup := e.Group(route.Path, middleware.Proxy(balancer))
		if route.Auth.User != "" {
			route.authorize(ctx, e, balancer)
		} else {
			e.Group(route.Path, middleware.Proxy(balancer))
		}

		if route.Rewrite {
			// rewrite URL! regular expression extensions to find only beginning of URL.
			// TODO: Maybe we can build a more flexible rewrite engine. This is just a simple solution!
			//rewritePaths["^"+route.Path+"/*"] = "/$1"
			rewritePaths[route.Path+"/*"] = "/$1"
		}

	}
	// eleminates path-parts like /auth /qr etc.
	e.Use(mcmiddleware.Rewrite(rewritePaths))
}
