package database

import (
	"github.com/jinzhu/gorm"
)

// ConfigStruct defines the current database connection
type ConfigStruct struct {
	Name     string `env:"DB_NAME" json:"name"`
	Server   string `env:"DB_HOST" json:"server"`
	User     string `env:"DB_USER" json:"user"`
	Password string `env:"DB_PASSWORD" json:"password"`
}

// DB is Global Database-Connection to Auth-Database
var DB *gorm.DB

// InitDatabase opens database-connections and do the automigration
func InitDatabase(dbConfig ConfigStruct) error {
	// Initialize DB-Connection
	// open database connection
	var err error
	DB, err = gorm.Open("mysql", getDatabasePath(dbConfig))
	if err != nil {
		DB.Close()
		return err
	}

	return nil
}

// getDatabasePath builds the database-string for Gorm. It will be built from
// the config parameters
func getDatabasePath(dbConfig ConfigStruct) string {
	return dbConfig.User + ":" + dbConfig.Password +
		"@(" + dbConfig.Server + ")/" + dbConfig.Name +
		"?charset=utf8&parseTime=True&loc=Local"
}
