package checkin

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type (
	// Checkin represents a single checkin of a user at a business-location. we save all data in a separate database
	// with any dependencies to other databases for user- or business-data. generally the stored checkin-data never
	// won't be updated expect the checkout-timestamp.
	Checkin struct {
		gorm.Model      `json:"-"`
		UUID            string    `json:"uuid" gorm:"type:varchar(36);unique_index"`
		BusinessUUID    string    `json:"-" gorm:"type:varchar(36)"`
		BusinessName    string    `json:"businessName" gorm:"type:varchar(50)"`
		BusinessAddress string    `json:"businessAddress" gorm:"type:varchar(300)"`
		UserUUID        string    `json:"-" gorm:"type:varchar(36)"`
		UserName        string    `json:"username" gorm:"type:varchar(500)"`
		UserPhone       string    `json:"userphone" gorm:"type:varchar(100)"`
		UserEmail       string    `json:"useremail" gorm:"type:varchar(255)"`
		UserStreet      string    `json:"userstreet" gorm:"type:varchar(500)"`
		UserCity        string    `json:"usercity" gorm:"type:varchar(100)"`
		UserCountry     string    `json:"usercountry" gorm:"type:varchar(100)"`
		UserRegistered  bool      `json:"-" gorm:"type:boolean;default:false"`
		Paper           bool      `json:"-" gorm:"type:boolean;default:false"`
		CheckedInAt     time.Time `json:"checkedInAt"`
		CheckedOutAt    time.Time `json:"checkedOutAt"`
	}
)

// BeforeCreate is a hook to set the UUID of a business at creating a new record
func (m *Checkin) BeforeCreate(tx *gorm.DB) (err error) {
	m.UUID = uuid.New().String()
	return
}
