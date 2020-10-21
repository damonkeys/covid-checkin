package main

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/damonkeys/ch3ck1n/monkeys/tracing"
	"github.com/labstack/echo/v4"
)

// readUserDataFromCookie read outs the user-data-json-string from cookie "user" and unmarshals it to a
// userData-struct.
func readUserDataFromCookie(e echo.Context) (*userData, error) {
	span := tracing.Enter(e)
	defer span.Finish()

	userData := userData{}
	cookie, err := e.Cookie("user")
	if err != nil {
		tracing.LogError(span, err)
		return nil, err
	}
	tracing.LogString(span, "user-cookie-value", cookie.Value)
	userDataUnescaped, err := url.QueryUnescape(cookie.Value)
	tracing.LogString(span, "user-cookie-value-unescaped", userDataUnescaped)
	// unmarshal JSON from cookie
	err = json.Unmarshal([]byte(userDataUnescaped), &userData)
	if err != nil {
		tracing.LogError(span, err)
		return nil, err
	}
	tracing.LogStruct(span, "user-struct", userData)
	return &userData, nil
}

// storeUserDataToCookie creates or updates cookie with given userData
func storeUserDataToCookie(e echo.Context, userData *userData) error {
	span := tracing.Enter(e)
	defer span.Finish()

	userDataJSON, err := json.Marshal(userData)
	if err != nil {
		tracing.LogError(span, err)
		return err
	}
	cookie := new(http.Cookie)
	cookie.Name = "user"
	cookie.Value = url.QueryEscape(string(userDataJSON))
	cookie.Domain = serverConfig.DomainName
	cookie.Secure = true
	e.SetCookie(cookie)
	tracing.LogString(span, "string set to cookie user", string(userDataJSON))
	return nil
}
