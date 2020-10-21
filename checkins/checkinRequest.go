package main

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/damonkeys/ch3ck1n/checkins/checkin"
	"github.com/damonkeys/ch3ck1n/monkeys/tracing"
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
		CheckedInAt:     time.Now(),
	}
	tracing.LogStruct(span, "checkin-model", checkin)
	// get User-UUID from session-cookie, if user is logged in
	sess, err := session.Get("_ch3ck1n_session", e)
	if err == nil {
		if sess.Values["userid"] != nil {
			tracing.LogString(span, "Session", "User logged in, session available.")
			checkin.UserUUID = sess.Values["userid"].(string)
			tracing.LogString(span, "User-ID", checkin.UserUUID)
		}
	}
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
	// store user-data from request in cookie for later checkins
	err = storeUserDataToCookie(e, checkinRequest.User)
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
