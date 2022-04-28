package main

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/damonkeys/covid-checkin/checkins/checkin"
	"github.com/damonkeys/covid-checkin/monkeys/tracing"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type (
	checkinRequest struct {
		User     *userData     `json:"user"`
		Business *businessData `json:"business"`
	}

	userData struct {
		Name    string `json:"name,required"`
		Street  string `json:"street,required"`
		City    string `json:"city,required"`
		Country string `json:"country,required"`
		Email   string `json:"email,required"`
		Phone   string `json:"phone,required"`
	}

	businessData struct {
		Name    string `json:"name,required"`
		Address string `json:"address"`
		UUID    string `json:"uuid,required"`
	}
)

// buildCheckinModel returns a checkin-model from all request and cookie-data.
func buildCheckinModel(e echo.Context) (*checkin.Checkin, error) {
	span := tracing.Enter(e)
	defer span.Finish()

	checkinRequest, err := getCheckinRequestData(e)
	if err != nil {
		tracing.LogError(span, err)
		return nil, err
	}

	tracing.LogStruct(span, "checkin-request", checkinRequest)
	checkin := &checkin.Checkin{
		BusinessName:    checkinRequest.Business.Name,
		BusinessAddress: checkinRequest.Business.Address,
		BusinessUUID:    checkinRequest.Business.UUID,
		UserName:        checkinRequest.User.Name,
		UserPhone:       checkinRequest.User.Phone,
		UserEmail:       checkinRequest.User.Email,
		UserStreet:      checkinRequest.User.Street,
		UserCity:        checkinRequest.User.City,
		UserCountry:     checkinRequest.User.Country,
		Paper:           false,
		UserRegistered:  false,
		CheckedInAt:     time.Now(),
	}
	tracing.LogStruct(span, "checkin-model", checkin)

	lastCheckinCookie, err := readUserDataFromCookie(e)
	// if no cookie we are happy and ignore that fact
	if err == nil {
		// if we find a cookie from last checkin we apply that userUUID to the new checkin for consistency
		checkin.UserUUID = lastCheckinCookie.UserUUID
	}

	sess, err := session.Get("_chckr_session", e)
	// if no session we are happy and ignore that fact (means the user hasn't logged in via authx yet)
	if err == nil {
		if sess.Values["userid"] != nil {
			tracing.LogString(span, "Session", "User logged in, session available. I set a flag for this in the checkin data")
			authxUserUUID := sess.Values["userid"].(string)
			if authxUserUUID != "" {
				checkin.UserRegistered = true
			}
		}
	}

	// store user-data from request in cookie for later checkins
	err = storeUserDataToCookie(e, checkin)
	if err != nil {
		tracing.LogError(span, err)
		return nil, err
	}
	tracing.LogString(span, "final User-UUID", checkin.UserUUID)
	return checkin, nil
}

// getCheckinRequestData returns a checkin-request-struct filled with data from the request.
// User-Data will be stored in user-cookie.
func getCheckinRequestData(e echo.Context) (*checkinRequest, error) {
	span := tracing.Enter(e)
	defer span.Finish()

	checkinRequest := &checkinRequest{}
	if err := e.Bind(checkinRequest); err != nil {
		tracing.LogError(span, err)
		return nil, err
	}
	tracing.LogStruct(span, "request", checkinRequest)
	err := validateCheckinRequest(tracing.GetContext(e), checkinRequest)
	if err != nil {
		tracing.LogError(span, err)
		return nil, err
	}

	return checkinRequest, nil
}

// validateCheckinRequest validates User and Business substruct of request based on set required tag.
func validateCheckinRequest(ctx context.Context, checkinRequest *checkinRequest) error {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	if err := checkRequiredFields(ctx, *checkinRequest.User); err != nil {
		return err
	}
	if err := checkRequiredFields(ctx, *checkinRequest.Business); err != nil {
		return err
	}
	return nil
}

// checkRequiredFields validates a struct based on additional required tag. If it is set, the field-value
// in the struct must not be empty.
func checkRequiredFields(ctx context.Context, s interface{}) error {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	tUser := reflect.TypeOf(s)
	for i := 0; i < tUser.NumField(); i++ {
		field := tUser.Field(i)
		tag := field.Tag.Get("json")
		if strings.Contains(tag, "required") {
			tUserVal := reflect.ValueOf(s)
			value := reflect.Indirect(tUserVal).FieldByName(field.Name)
			if value.String() == "" {
				return fmt.Errorf("%s is missing in request", field.Name)
			}
		}
	}
	return nil
}
