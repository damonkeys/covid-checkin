package main

// ch3ck1nweb webserver for static websites. To serve your static site put all files in folder static.
//
// The server uses environment variables. If they are not set the server won't start. It expects the following environment variables:
//   * SERVER_PORT - the server is listening on this portnumber

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/damonkeys/covid-checkin/monkeys/config"
	l "github.com/damonkeys/covid-checkin/monkeys/logger"
	"github.com/damonkeys/covid-checkin/monkeys/tracing"
)

type (
	// ServerConfigStruct holds the server-config
	ServerConfigStruct struct {
		Port string `env:"SERVER_PORT"`
	}
)

var serverConfig ServerConfigStruct

func main() {
	// tracer init
	closer, span, ctx := tracing.InitJaeger("ch3ck1nweb")
	defer closer.Close()

	// Initialize echo and set logger
	e := echo.New()
	e.Use(tracing.MiddlewareWithoutCurrentUser("ch3ck1nweb"))
	l.ConfigureLogger(ctx, "ch3ck1nweb", e)

	// read config from environment variables to struct
	configInterface, err := config.ReadEnvVars(ctx, ServerConfigStruct{})
	if err != nil {
		e.Logger.Error(err)
		tracing.LogError(span, err)
		span.Finish()
		os.Exit(-1)
	}
	serverConfig = configInterface.(ServerConfigStruct)

	// Start webserver
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static",
		Browse: true,
		HTML5:  true,
	}))
	e.Use(middleware.Recover())
	span.Finish()
	e.Logger.Fatal(e.Start(":" + serverConfig.Port))
}
