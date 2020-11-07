package main

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/damonkeys/ch3ck1n/checkins/checkin"
	"github.com/damonkeys/ch3ck1n/monkeys/tracing"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const cookiename = "checkin"

// readUserDataFromCookie read outs the user-data-json-string from cookie "user" and unmarshals it to a
// checkin.Checkin struct.
func readUserDataFromCookie(e echo.Context) (*checkin.Checkin, error) {
	span := tracing.Enter(e)
	defer span.Finish()

	checkin := &checkin.Checkin{}
	cookie, err := e.Cookie(cookiename)
	if err != nil {
		tracing.LogError(span, err)
		return nil, err
	}
	tracing.LogString(span, "checkin-cookie-value", cookie.Value)
	checkinDataUnescaped, err := url.QueryUnescape(cookie.Value)
	tracing.LogString(span, "checkin-cookie-value-unescaped", checkinDataUnescaped)
	// unmarshal JSON from cookie
	err = json.Unmarshal([]byte(checkinDataUnescaped), checkin)
	if err != nil {
		tracing.LogError(span, err)
		return nil, err
	}
	tracing.LogStruct(span, "checkin-struct", checkin)
	return checkin, nil
}

// storeUserDataToCookie creates or updates cookie with given checkin.Checkin
// One of the important functions is that it establishes a uuid for the user checkin data
// no matter wether the user is authenticated via authx or
// not (meaning the uuid is generated)
func storeUserDataToCookie(e echo.Context, checkin *checkin.Checkin) error {

	span := tracing.Enter(e)
	defer span.Finish()

	if checkin.UserUUID == "" { //its prefilled if session cookie from authx exist
		checkin.UserUUID = uuid.New().String() // set a cookie based uuid to detect revisits from user
	}
	userDataJSON, err := json.Marshal(checkin)
	if err != nil {
		tracing.LogError(span, err)
		return err
	}
	cookie := &http.Cookie{
		Name:     cookiename,
		Value:    url.QueryEscape(string(userDataJSON)),
		Domain:   serverConfig.DomainName,
		Secure:   true,
		MaxAge:   2147483648, // https://stackoverflow.com/questions/3290424/set-a-cookie-to-never-expire#3290474
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
	}
	e.SetCookie(cookie)
	tracing.LogString(span, "string set to cookie user", string(userDataJSON))
	return nil
}
