package tracing

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// Middleware defines the tracing middleware. Use it with echo as first middleware for tracing the whole request time!
func Middleware(servername string) echo.MiddlewareFunc {
	return func(nextHandlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Before Request - tracer init
			closer, span, ctx := InitJaeger(servername, c.Request().URL.Path)
			c.Set("tracingctx", ctx)
			span.SetTag("request-path", c.Request().URL)

			// get session user-uuid - we don't need to check valid session, we do it in the middleware
			sess, _ := session.Get("_chckr_session", c)
			currentUserUUID := sess.Values["userid"]
			if currentUserUUID != nil {
				span.SetTag("current-user-uuid", currentUserUUID.(string))
			}

			f := nextHandlerFunc(c)

			// After request
			span.Finish()
			closer.Close()
			return f
		}
	}
}

// MiddlewareWithoutCurrentUser defines the tracing middleware without logging the currentUserUUID. Use it with echo as first middleware for tracing the whole request time!
func MiddlewareWithoutCurrentUser(servername string) echo.MiddlewareFunc {
	return func(nextHandlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Before Request - tracer init
			closer, span, ctx := InitJaeger(servername, c.Request().URL.Path)
			c.Set("tracingctx", ctx)
			span.SetTag("request-path", c.Request().URL)

			f := nextHandlerFunc(c)

			// After request
			span.Finish()
			closer.Close()
			return f
		}
	}
}
