package main

// # business - Server for all known business in chckr
//
// ## The server uses environment variables. If they are not set the server won't start. It expects the following environment variables:
//   * SERVER_PORT       - the server is listening on this portnumber
//   * DB_HOST           - database host for connecting the auth database
//   * DB_NAME           - database name for connecting the auth database
//   * DB_USER           - database user for connecting the auth database
//   * DB_PASSWORD       - database user-password for connecting the auth database
import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/markbates/goth/gothic"

	"github.com/damonkeys/covid-checkin/biz/business"
	"github.com/damonkeys/covid-checkin/monkeys/config"
	"github.com/damonkeys/covid-checkin/monkeys/database"
	l "github.com/damonkeys/covid-checkin/monkeys/logger"
	"github.com/damonkeys/covid-checkin/monkeys/tracing"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
	// ServerConfigStruct holds the server-config for biz
	ServerConfigStruct struct {
		Port          string                `env:"SERVER_PORT"`
		SessionSecret string                `env:"SESSION_SECRET"`
		Database      database.ConfigStruct `json:"database"`
	}
)

const sessionName = "_chckr_callback"
const serverName = "biz"

// serverConfig defines the configuration for auth
var serverConfig ServerConfigStruct

func main() {
	// tracer init
	closer, span, ctx := tracing.InitJaeger(serverName)
	defer closer.Close()

	// Init echo
	e := echo.New()
	l.ConfigureLogger(ctx, serverName, e)
	readEnvironmentConfig(ctx, e.Logger)

	if err := database.InitDatabase(serverConfig.Database); err != nil {
		e.Logger.Fatal(err)
		tracing.LogError(span, err)
		span.Finish()
		os.Exit(0)
	}
	defer database.DB.Close()

	// creeate session store for echo and gorilla (used by goth!)
	sessionStore := sessions.NewCookieStore([]byte(serverConfig.SessionSecret))
	gothic.Store = sessionStore
	e.Use(session.Middleware(sessionStore))
	e.Use(tracing.Middleware(serverName))
	e.Use(middleware.Recover())

	// Routes
	e.GET("/business/:code", getLocation)

	span.Finish()
	e.Logger.Fatal(e.Start(":" + serverConfig.Port))
}

// readEnvironmentConfig reads all needed environment variables and save it in serverConfig struct
func readEnvironmentConfig(ctx context.Context, log echo.Logger) {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	// read config from environment variables to struct
	configInterface, err := config.ReadEnvVars(ctx, ServerConfigStruct{})
	if err != nil {
		log.Error(err)
		tracing.LogError(span, err)
		os.Exit(-1)
	}
	serverConfig = configInterface.(ServerConfigStruct)
}

// business returns all business-data for the given business-code
func getLocation(e echo.Context) error {
	span := tracing.Enter(e)
	defer span.Finish()

	locationCode := e.Param("code")
	if locationCode == "" {
		err := errors.New("no business-code received")
		tracing.LogError(span, err)
		return serverErrorOnError(err, e)
	}
	span.SetTag("business-code", locationCode)
	tracing.LogString(span, "log", fmt.Sprintf("Get business-data with code %s", locationCode))

	ctx := tracing.GetContext(e)

	businessOperation := &business.Operations{
		Business: business.Business{},
		Context:  ctx,
	}

	err := businessOperation.GetBusinessByCode(locationCode)
	if err != nil {
		tracing.LogError(span, err)
		return serverErrorOnError(err, e)
	}
	return e.JSON(http.StatusOK, businessOperation.Business)
}

// serverErrorOnError returns an internalServerError for a given err
func serverErrorOnError(err error, e echo.Context) error {
	span := tracing.Enter(e)
	defer span.Finish()

	if err != nil {
		tracing.LogError(span, err)
		e.Logger().Warnf("serverErrorOnError error: %s", err)
		return e.JSON(http.StatusInternalServerError, "error")
	}
	return nil
}
