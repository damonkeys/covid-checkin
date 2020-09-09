package models

import (
	"context"
	"errors"
	"time"

	"github.com/damonkeys/ch3ck1n/monkeys/database"
	"github.com/damonkeys/ch3ck1n/monkeys/tracing"
	"github.com/jinzhu/gorm"
)

type (
	// Provider holds all data of a OAuth-Provider of a user
	Provider struct {
		gorm.Model
		ProviderName string `gorm:"type:varchar(50)"`
		AccessToken  string `gorm:"type:varchar(255)"`
		RefreshToken string `gorm:"type:varchar(255)"`
		Name         string `gorm:"type:varchar(255);"`
		Lastname     string `gorm:"type:varchar(255);"`
		Firstname    string `gorm:"type:varchar(255);"`
		UserID       string `gorm:"type:varchar(255);"`
		AvatarURL    string `gorm:"type:varchar(255);"`
		Nickname     string `gorm:"type:varchar(255);"`
		Location     string `gorm:"type:varchar(255);"`
		ExpiresAt    time.Time
		UserRefer    uint
	}
)

// Create inserts a new provider record in database
func (provider *Provider) Create(ctx context.Context) error {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	result := database.DB.Create(&provider)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete (soft-)deletes the given provider
func (provider *Provider) Delete(ctx context.Context, db *gorm.DB) error {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	if db == nil {
		db = database.DB
	}
	result := db.Delete(&provider)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Undelete removes deletedAt field
func (provider *Provider) Undelete(ctx context.Context) error {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	error := database.DB.Model(&provider).Unscoped().Select("deleted_at").Updates(map[string]interface{}{"deleted_at": nil})
	return error.Error
}

// FindUserProviderByName returns a provider of the given user if it exists.
func FindUserProviderByName(ctx context.Context, user User, providerName string) (*Provider, error) {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	database.DB.Model(&user).Association("Providers").Find(&user.Providers)
	for _, provider := range user.Providers {
		if provider.ProviderName == providerName {
			return &provider, nil
		}
	}
	return new(Provider), errors.New("no existing provider found")
}

// FindProvidersByUser returns all providers assiciated with user
func FindProvidersByUser(ctx context.Context, user *User, db *gorm.DB) []Provider {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	if db == nil {
		db = database.DB
	}
	db.Model(&user).Association("Providers").Find(&user.Providers)
	return user.Providers
}

// FindDeletedProvidersByUser search given email address in provider table and returns soft-deleted provider
func FindDeletedProvidersByUser(ctx context.Context, user *User) []Provider {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	database.DB.Unscoped().Model(&user).Association("Providers").Find(&user.Providers)
	return user.Providers
}

// DeleteAllProvidersByUser deletes all providers by given user
func DeleteAllProvidersByUser(ctx context.Context, user *User, db *gorm.DB) error {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	providers := FindProvidersByUser(ctx, user, db)
	for _, provider := range providers {
		err := provider.Delete(ctx, db)
		if err != nil {
			return err
		}
	}
	return nil
}

// UndeleteAllProvidersByUser deletes all providers by given user
func UndeleteAllProvidersByUser(ctx context.Context, user *User) error {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	providers := FindDeletedProvidersByUser(ctx, user)
	for _, provider := range providers {
		err := provider.Undelete(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}
