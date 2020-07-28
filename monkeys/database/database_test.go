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

func TestInitDatabase(t *testing.T) {
	databaseConfig := ConfigStruct{
		Name:     "test_monkey_auth",
		Password: "_not_valid",
		User:     "auth_user",
		Server:   "",
	}
	// global variable DB might be nil
	if DB != nil {
		t.Error("global variable is not nil without calling InitDatabase")
	}
	// test with invalid password
	err := InitDatabase(databaseConfig)
	if err == nil {
		t.Error("database connection doesn't fail - password might be wrong!")
	}
	// test with valid password
	databaseConfig.Password = ""
	err = InitDatabase(databaseConfig)
	if err != nil {
		t.Errorf("database connection failed: %s", err)
	}
	// global variable DB is set
	if DB == nil {
		t.Error("global variable DB for db-connections is nil")
	}
}
