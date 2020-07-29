package logging

// Common Logger-Settings. Maybe we do it configurable?
import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/damonkeys/ch3ck1n/monkeys/tracing"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

// Log-Config... maybe a todo to make it configurable
const logLevel = log.DEBUG
const logPath = "../logs/"

// ANSI colors
const white = "\033[37m"
const red = "\033[31m"
const green = "\033[32m"
const yellow = "\033[33m"
const blue = "\033[94m"
const lightwhite = "\033[97m"
const cyan = "\033[36m"

// ConfigureLogger do a consitient logger-configuration for all microservices
func ConfigureLogger(ctx context.Context, name string, e *echo.Echo) {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	err := os.MkdirAll(logPath, 0755)
	if err != nil {
		fmt.Printf("Something went wrong during log directory initialisation: %s", err)
	}
	logfileWriter, err := os.OpenFile(logPath+name+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("ERROR: Cannot create log-file: %s", err)
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: io.MultiWriter(os.Stdout, logfileWriter),
		Format: white + "${time_rfc3339}  " + green + "${method}" + white + "\tfrom " + blue + "${remote_ip}" + white +
			"\twith status " + yellow + "${status}" + white + "\t" + cyan +
			"latency: ${latency}\tlatency_human: ${latency_human}\tbytes_in: ${bytes_in}\tbytes_out: ${bytes_out}\tid: ${id}\t" + blue +
			"${host}${uri}" + white + "\t" + red + "${error}\n" + white,
	}))
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader(white + "${time_rfc3339}  " + green + "${level}" + white + "\t" + blue + "${short_file}" + yellow + ":${line}\t" + lightwhite)
		l.SetOutput(io.MultiWriter(os.Stdout, logfileWriter))
	}
	e.Logger.SetLevel(logLevel)

	e.Logger.Info(red + "\n-----------------" + white)
}
