package business

import (
	"context"
	"fmt"

	"github.com/damonkeys/ch3ck1n/monkeys/database"
	"github.com/damonkeys/ch3ck1n/monkeys/tracing"
)

type (
	// Finder finds a business by a given code
	Finder interface {
		GetBusinessByCode(code string) error
	}

	// Operations is the struct that enables manupulations on businesses such as CRUD and Finder operations
	Operations struct {
		database.CRUD
		Business Business
		Context  context.Context
	}
)

// GetBusinessByCode returns a business record fot the given code
func (o *Operations) GetBusinessByCode(code string) error {
	span := tracing.EnterWithContext(o.Context)
	defer span.Finish()
	tracing.LogString(span, "code", code)

	dbResult := database.DB.Where("code = ?", code).Preload("BusinessInfos").Find(&o.Business)
	if dbResult.Error != nil {
		tracing.LogError(span, fmt.Errorf("Cannot find business with code: %s", code))
		return dbResult.Error
	}
	tracing.LogStruct(span, "foundLBusinessInDB", &o.Business)
	return nil
}

// Create stores the whole Business record. The UUID will be created automatically.
func (o *Operations) Create() error {
	span := tracing.EnterWithContext(o.Context)
	defer span.Finish()
	tracing.LogStruct(span, "business", o.Business)
	result := database.DB.Create(o.Business)
	if result.Error != nil {
		tracing.LogError(span, result.Error)
		return result.Error
	}
	return nil
}

// Update updates the whole Business record. It isn't possible to change UUID!
func (o *Operations) Update() error {
	span := tracing.EnterWithContext(o.Context)
	defer span.Finish()
	tracing.LogStruct(span, "business", o.Business)
	result := database.DB.Save(o.Business)
	if result.Error != nil {
		tracing.LogError(span, result.Error)
		return result.Error
	}
	return nil
}

// Delete (soft-)deletes the given Business
func (o *Operations) Delete() error {
	span := tracing.EnterWithContext(o.Context)
	defer span.Finish()
	tracing.LogStruct(span, "business", o.Business)

	result := database.DB.Delete(o.Business)
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
	tracing.LogStruct(span, "business", o.Business)

	error := database.DB.Model(o.Business).Unscoped().Select("deleted_at").Updates(map[string]interface{}{"deleted_at": nil})
	return error.Error
}
