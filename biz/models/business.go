package models

import (
	"github.com/jinzhu/gorm"
)

type (
	// Business represents a single location with all its address-data
	Business struct {
		gorm.Model    `json:"-"`
		UUID          string `json:"uuid" gorm:"type:varchar(36);unique_index"`
		Name          string `json:"name" gorm:"type:varchar(15)"`
		Code          string `json:"code" gorm:"type:varchar(5);unique_index"`
		BusinessInfos []BusinessInfo
	}

	// BusinessInfo represents detailed description for a location in different languages
	BusinessInfo struct {
		gorm.Model  `json:"-"`
		UUID        string `json:"uuid" gorm:"type:varchar(36);unique_index"`
		Description string `json:"description"`
		Language    string `json:"language"`
		BusinessID  uint   `json:"-"`
	}
)
