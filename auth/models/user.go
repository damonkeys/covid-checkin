package models

import (
	"context"
	"fmt"
	"time"

	"../../monkeys/database"
	"../../monkeys/tracing"
	"github.com/jinzhu/gorm"
)

type (
	// User holds all data of a user in a database
	User struct {
		gorm.Model
		UUID                    string     `gorm:"type:varchar(36);unique_index"`
		Email                   string     `gorm:"type:varchar(255);unique_index"`
		Name                    string     `gorm:"type:varchar(255)"`
		AvatarURL               string     `gorm:"type:varchar(512)"`
		Merchant                bool       `gorm:"type:boolean"`
		Providers               []Provider `gorm:"foreignkey:UserRefer;association_foreignkey:id"`
		Active                  bool       `gorm:"type:boolean"`
		ActivationToken         string     `gorm:"type:varchar(36)"`
		ActivationTokenCreation time.Time
		ActiveSince             time.Time
	}

	// UserInterface defines all functions for User-Model
	UserInterface interface {
		database.Model
		AppendProviderToUser(ctx context.Context, provider *Provider) error
		AfterDelete(tx *gorm.DB) error
	}
)

// FindUserByEmail search given email address in user table
func FindUserByEmail(ctx context.Context, email string) (*User, error) {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	tracing.LogString(span, "EMail", email)
	authUser := new(User)
	database.DB.Where("email = ?", email).First(&authUser)
	tracing.LogStruct(span, "AuthUser after search by Email", authUser)
	if authUser.Email == "" {
		tracing.LogError(span, fmt.Errorf("no existing User with email-address %s found", email))
		return authUser, fmt.Errorf("no existing User with email-address %s found", email)
	}
	tracing.LogString(span, "log", "Looking for providers")
	database.DB.Model(&authUser).Association("Providers").Find(&authUser.Providers)
	return authUser, nil
}

// FindDeletedUserByEmail search given email address in user table and returns soft-deleted user
func FindDeletedUserByEmail(ctx context.Context, email string) (*User, error) {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	authUser := new(User)
	database.DB.Unscoped().Where("email = ?", email).First(&authUser)
	if authUser.Email == "" || authUser.DeletedAt == nil {
		return authUser, fmt.Errorf("no soft-deleted User with email-address %s found", email)
	}
	database.DB.Model(&authUser).Association("Providers").Find(&authUser.Providers)
	return authUser, nil
}

// FindUserByUUID search given email address in user table
func FindUserByUUID(ctx context.Context, uuid string) (*User, error) {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	authUser := new(User)
	database.DB.Where("uuid = ?", uuid).First(&authUser)
	if authUser.Email == "" {
		return authUser, fmt.Errorf("no existing User with uuid %s found", uuid)
	}
	database.DB.Model(&authUser).Association("Providers").Find(&authUser.Providers)
	return authUser, nil
}

// Create inserts a new user record in database
func (user *User) Create(ctx context.Context) error {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	result := database.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Update updates the whole user record. It isn't possible to change UUID!
func (user *User) Update(ctx context.Context) error {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	_, err := FindUserByUUID(ctx, user.UUID)
	if err != nil {
		return err
	}
	result := database.DB.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete (softdeletes) user and all associated providers.
func (user *User) Delete(ctx context.Context) error {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	result := database.DB.Delete(&user)
	if result.Error != nil {
		tracing.LogError(span, result.Error)
		return result.Error
	}
	return nil
}

// Undelete removes deletedAt field
func (user *User) Undelete(ctx context.Context) error {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	error := database.DB.Model(&user).Unscoped().Select("deleted_at").Updates(map[string]interface{}{"deleted_at": nil})
	UndeleteAllProvidersByUser(ctx, user)
	return error.Error
}

// AppendProviderToUser adds a OAuth-provider to the user
func (user *User) AppendProviderToUser(ctx context.Context, provider *Provider) error {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	database.DB.Model(&user).Association("Providers")
	result := database.DB.Model(&user).Association("Providers").Append(provider)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Hooks

// AfterDelete defines hook for after deleting a user. You have to delete all providers, too.
func (user *User) AfterDelete(tx *gorm.DB) (err error) {
	// TODO tracing
	_, _, ctx := tracing.InitMockJaeger("models-test")
	return DeleteAllProvidersByUser(ctx, user, tx)
}
