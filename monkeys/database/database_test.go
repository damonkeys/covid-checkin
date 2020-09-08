package database

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestGetDatabasePath(t *testing.T) {
	databaseConfig := ConfigStruct{
		Name:     "name",
		Password: "password",
		User:     "username",
		Server:   "server",
	}
	databasePath := getDatabasePath(databaseConfig)

	if databasePath != "username:password@server/name?charset=utf8&parseTime=True&loc=Local" {
		t.Error("wrong databasepath will be returned")
	}
}
