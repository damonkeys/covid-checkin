package checkin

import (
	"context"

	"github.com/damonkeys/ch3ck1n/monkeys/database"
	"github.com/damonkeys/ch3ck1n/monkeys/tracing"
)

type (
	// Operations is the struct that holds all possible opertaions on the business-data
	Operations struct {
		database.CRUD
		Context     context.Context
		CheckinData *Checkin
	}
)

// Create stores the checkin record. The UUID will be created automatically.
func (o *Operations) Create() error {
	span := tracing.EnterWithContext(o.Context)
	defer span.Finish()
	tracing.LogStruct(span, "checkin", o.CheckinData)
	result := database.DB.Create(o.CheckinData)
	if result.Error != nil {
		tracing.LogError(span, result.Error)
		return result.Error
	}
	return nil
}

// Update updates the checkin record.
func (o *Operations) Update() error {
	span := tracing.EnterWithContext(o.Context)
	defer span.Finish()
	tracing.LogStruct(span, "checkin", o.CheckinData)
	result := database.DB.Save(o.CheckinData)
	if result.Error != nil {
		tracing.LogError(span, result.Error)
		return result.Error
	}
	return nil
}

// Delete (soft-)deletes the given Checkin
func (o *Operations) Delete() error {
	span := tracing.EnterWithContext(o.Context)
	defer span.Finish()
	tracing.LogStruct(span, "checkin", o.CheckinData)

	result := database.DB.Delete(o.CheckinData)
	if result.Error != nil {
		tracing.LogError(span, result.Error)
		return result.Error
	}
	return nil
}

// Undelete removes deletedAt field
func (o *Operations) Undelete() error {
	span := tracing.EnterWithContext(o.Context)
	defer span.Finish()
	tracing.LogStruct(span, "checkin", o.CheckinData)

	result := database.DB.Model(o.CheckinData).Unscoped().Select("deleted_at").Updates(map[string]interface{}{"deleted_at": nil})
	if result.Error != nil {
		tracing.LogError(span, result.Error)
		return result.Error
	}
	return nil
}