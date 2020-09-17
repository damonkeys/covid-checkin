package business

import (
	"os"

	"github.com/damonkeys/ch3ck1n/biz/qr"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type (
	// Business represents a single location with all its address-data
	Business struct {
		gorm.Model    `json:"-"`
		UUID          string `json:"uuid" gorm:"type:varchar(36);unique_index"`
		Name          string `json:"name" gorm:"type:varchar(50)"`
		Code          string `json:"code" gorm:"type:varchar(5);unique_index"`
		Street1       string `json:"street1" gorm:"type:varchar(50)"`
		Street2       string `json:"street2" gorm:"type:varchar(50)"`
		Zip           string `json:"zip" gorm:"type:varchar(10)"`
		City          string `json:"city" gorm:"type:varchar(30)"`
		Country       string `json:"country" gorm:"type:varchar(30)"`
		BusinessInfos []BusinessInfo
	}

	// BusinessInfo represents detailed description for a location in different languages
	BusinessInfo struct {
		gorm.Model  `json:"-"`
		UUID        string `json:"uuid" gorm:"type:varchar(36);unique_index"`
		Description string `json:"description" gorm:"type:mediumtext"`
		Language    string `json:"language"`
		BusinessID  uint   `json:"-"`
	}
)

// BeforeCreate is a hook to set the UUID of a business at creating a new record
func (b *Business) BeforeCreate(tx *gorm.DB) (err error) {
	b.UUID = uuid.New().String()
	code, err := generateCode()
	if err != nil {
		return err
	}
	b.Code = code
	callToPixi := qr.NewPixiCall()
	fh := callToPixi.Handler.(*qr.FileHandler)
	deeplink := os.Getenv("DEEP_LINK_TO_BUSINESS_BY_CODE")
	fh.Filename = b.UUID + ".png"
	err = callToPixi.Retrieve(deeplink+code, "", "png")
	if err != nil {
		return err
	}
	fh.Filename = b.UUID + ".svg"
	err = callToPixi.Retrieve(deeplink+code, "", "svg")
	if err != nil {
		return err
	}
	return nil

}

// BeforeCreate is a hook to set the UUID of a business at creating a new record
func (bi *BusinessInfo) BeforeCreate(tx *gorm.DB) (err error) {
	bi.UUID = uuid.New().String()
	return
}
