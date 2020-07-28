package middleware

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	// RewriteConfig defines the config for Rewrite middleware.
	RewriteConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper middleware.Skipper

		// Rules defines the URL path rewrite rules. The values captured in asterisk can be
		// retrieved by index e.g. $1, $2 and so on.
		// Example:
		// "/old":              "/new",
		// "/api/*":            "/$1",
		// "/js/*":             "/public/javascripts/$1",
		// "/users/*/orders/*": "/user/$1/order/$2",
		// Required.
		Rules map[string]string `yaml:"rules"`

		rulesRegex map[*regexp.Regexp]string
	}
)

var (
	// DefaultRewriteConfig is the default Rewrite middleware config.
	DefaultRewriteConfig = RewriteConfig{
		Skipper: middleware.DefaultSkipper,
	}
)

// Rewrite returns a Rewrite middleware.
//
// Rewrite middleware rewrites the URL path based on the provided rules.
func Rewrite(rules map[string]string) echo.MiddlewareFunc {
	c := DefaultRewriteConfig
	c.Rules = rules
	return RewriteWithConfig(c)
}

// RewriteWithConfig returns a Rewrite middleware with config.
// See: `Rewrite()`.
func RewriteWithConfig(config RewriteConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Rules == nil {
		panic("echo: rewrite middleware requires url path rewrite rules")
	}
	if config.Skipper == nil {
		config.Skipper = middleware.DefaultBodyDumpConfig.Skipper
	}
	config.rulesRegex = map[*regexp.Regexp]string{}

	// Initialize
	for k, v := range config.Rules {
		k = regexp.QuoteMeta(k)
		k = strings.Replace(k, `\*`, "(.*)", -1)
		k = k + "$"
		config.rulesRegex[regexp.MustCompile(k)] = v
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			useDefaultRoute := true
			// Rewrite
			for k, v := range config.rulesRegex {
				// Default route?
				if k.String() != "/(.*)$" &&
					k.String() != "//(.*)$" &&
					k.String() != "/(.*)/(.*)$" {
					replacer := captureTokens(k, req.URL.Path)
					if replacer != nil {
						req.URL.Path = replacer.Replace(v)
						useDefaultRoute = false
						break
					}
				}
			}
			// No real route found? Use Default-Routing for /*
			// This is a quick hack for using static sites with echo rewrite middleware
			if useDefaultRoute {
				k := regexp.QuoteMeta("/*")
				k = strings.Replace(k, `\*`, "(.*)", -1)
				k = k + "$"
				replacer := captureTokens(regexp.MustCompile(k), req.URL.Path)
				if replacer != nil {
					req.URL.Path = replacer.Replace("/$1")
				}
			}
			return next(c)
		}
	}
}

// From middleware.go
func captureTokens(pattern *regexp.Regexp, input string) *strings.Replacer {
	groups := pattern.FindAllStringSubmatch(input, -1)
	if groups == nil {
		return nil
	}
	values := groups[0][1:]
	replace := make([]string, 2*len(values))
	for i, v := range values {
		j := 2 * i
		replace[j] = "$" + strconv.Itoa(i+1)
		replace[j+1] = v
	}
	return strings.NewReplacer(replace...)
}
