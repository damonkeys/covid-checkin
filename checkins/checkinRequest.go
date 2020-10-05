package main

import (
	"context"
	"time"

	"github.com/damonkeys/ch3ck1n/checkins/checkin"
	"github.com/damonkeys/ch3ck1n/monkeys/tracing"
)

type (
	checkinRequest struct {
		Chckr    chckrData    `json:"chckr"`
		Business businessData `json:"business"`
	}

	chckrData struct {
		Name    string `json:"name"`
		Street  string `json:"street"`
		City    string `json:"city"`
		Country string `json:"country"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
	}

	businessData struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		UUID    string `json:"uuid"`
	}
)

func buildCheckinModel(c context.Context, checkinRequest checkinRequest) *checkin.Checkin {
	span := tracing.EnterWithContext(c)
	defer span.Finish()
	tracing.LogStruct(span, "checkin-request", checkinRequest)

	checkin := &checkin.Checkin{
		BusinessName:    checkinRequest.Business.Name,
		BusinessAddress: checkinRequest.Business.Address,
		BusinessUUID:    checkinRequest.Business.UUID,
		ChckrName:       checkinRequest.Chckr.Name,
		ChckrPhone:      checkinRequest.Chckr.Phone,
		ChckrEmail:      checkinRequest.Chckr.Email,
		ChckrStreet:     checkinRequest.Chckr.Street,
		ChckrCity:       checkinRequest.Chckr.City,
		ChckrCountry:    checkinRequest.Chckr.Country,
		Paper:           false,
		CheckedInAt:     time.Now(),
	}
	return checkin
}
