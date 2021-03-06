package session

import (
	"github.com/damonkeys/covid-checkin/monkeys/tracing"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func UserUUIDFromSession(e echo.Context) string {
	span := tracing.Enter(e)
	defer span.Finish()

	// get session user-uuid - we don't need to check valid session, we do it in the middleware
	sess, _ := session.Get("_chckr_session", e)
	return sess.Values["userid"].(string)
}
