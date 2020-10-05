package main

// # checkins - Server for all known checkins in chckr
//
// ## The server uses environment variables. If they are not set the server won't start. It expects the following environment variables:
//   * SERVER_PORT       - the server is listening on this portnumber
//   * DB_HOST           - database host for connecting the checkin database
//   * DB_NAME           - database name for connecting the checkin database
//   * DB_USER           - database user for connecting the checkin database
//   * DB_PASSWORD       - database user-password for connecting the checkin database
//   * SESSION_SECRET    - The secret that we use for secure sessions
import (
	"context"
	"net/http"
	"os"

	"github.com/gorilla/sessions"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/markbates/goth/gothic"

	"github.com/damonkeys/ch3ck1n/checkins/checkin"
	"github.com/damonkeys/ch3ck1n/monkeys/config"
	"github.com/damonkeys/ch3ck1n/monkeys/database"
	l "github.com/damonkeys/ch3ck1n/monkeys/logger"
	"github.com/damonkeys/ch3ck1n/monkeys/tracing"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
	// ServerConfigStruct holds the server-config for checkin
	ServerConfigStruct struct {
		Port          string                `env:"SERVER_PORT"`
		SessionSecret string                `env:"SESSION_SECRET"`
		Database      database.ConfigStruct `json:"database"`
	}
)

const sessionName = "_ch3ck1n_callback"
const serverName = "checkins"

// serverConfig defines the configuration for checkin
var serverConfig ServerConfigStruct

func main() {
	// tracer init
	closer, span, ctx := tracing.InitJaeger(serverName)
	defer closer.Close()

	// Init echo
	e := echo.New()
	l.ConfigureLogger(ctx, serverName, e)
	readEnvironmentConfig(ctx, e.Logger)
	tracing.LogStruct(span, "serverConfig", serverConfig)

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
	e.POST("/checkin", checkinAtBusiness)

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

// checkinAtBusiness saves a checkin of a chckr to the checkin-database. The request contains the chckr- and the
// business-data. Checkin will store in database with timestamp for checkin of the request.
func checkinAtBusiness(e echo.Context) error {
	span := tracing.Enter(e)
	defer span.Finish()

	checkinRequest := &checkinRequest{}
	if err := e.Bind(checkinRequest); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad checkin")
	}
	tracing.LogStruct(span, "request", checkinRequest)

	checkinData := buildCheckinModel(tracing.GetContext(e), *checkinRequest)
	sess, err := session.Get("_ch3ck1n_session", e)
	if err == nil {
		if sess.Values["userid"] != nil {
			checkinData.ChckrUUID = sess.Values["userid"].(string)
		}
	}
	operations := &checkin.Operations{
		Context:     tracing.GetContext(e),
		CheckinData: checkinData,
	}

	err = operations.Create()
	if err != nil {
		tracing.LogError(span, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "checkin failed")
	}
	return e.String(http.StatusOK, "")
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
