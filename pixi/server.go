package main

// pixi webserver for static content. We use this server to host binary stuff we serve.
// Currently this is only qr codes.
//
// The server uses environment variables. If they are not set the server won't start. It expects the following environment variables:
//   * SERVER_PORT - the server is listening on this portnumber

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/damonkeys/ch3ck1n/monkeys/config"
	l "github.com/damonkeys/ch3ck1n/monkeys/logger"
	"github.com/damonkeys/ch3ck1n/monkeys/tracing"
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
	closer, span, ctx := tracing.InitJaeger("pixi")
	defer closer.Close()

	// Initialize echo and set logger
	e := echo.New()
	e.Use(tracing.MiddlewareWithoutCurrentUser("pixi"))
	l.ConfigureLogger(ctx, "pixi", e)

	// read config from environment variables to struct
	configInterface, err := config.ReadEnvVars(ctx, ServerConfigStruct{})
	if err != nil {
		e.Logger.Error(err)
		tracing.LogError(span, err)
		span.Finish()
		os.Exit(-1)
	}
	serverConfig = configInterface.(ServerConfigStruct)
	// Trace config
	tracing.LogStruct(span, "config", serverConfig)

	// Start webserver
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "static",
		Browse: false,
		HTML5:  true,
	}))
	e.Use(middleware.Recover())
	span.Finish()
	e.Logger.Fatal(e.Start(":" + serverConfig.Port))
}
