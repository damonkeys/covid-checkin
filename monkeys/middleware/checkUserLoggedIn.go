package middleware

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// CheckLoggedInUserAccess checks wether a user is logged in. If not the request ist not further processed and 403 is returned.
func CheckLoggedInUserAccess() echo.MiddlewareFunc {
	return func(nextHandlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(e echo.Context) error {
			sess, err := getSession(e)
			if err != nil {
				return e.JSON(http.StatusForbidden, "not-allowed")
			}
			if sess.Values["userid"] == nil {
				return e.JSON(http.StatusForbidden, "not-allowed")
			}
			return nextHandlerFunc(e)
		}
	}
}

func getSession(e echo.Context) (*sessions.Session, error) {
	return session.Get("_chckr_session", e)
}
