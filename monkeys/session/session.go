package session

import (
	"../tracing"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

func UserUUIDFromSession(e echo.Context) string {
	span := tracing.Enter(e)
	defer span.Finish()

	// get session user-uuid - we don't need to check valid session, we do it in the middleware
	sess, _ := session.Get("_monkeycash_session", e)
	return sess.Values["userid"].(string)
}
