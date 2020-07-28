package models

import (
	"testing"

	"../../monkeys/database"
	"github.com/google/uuid"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestCreateProvider(t *testing.T) {
	setupDatabaseTests()
	testProvider := &Provider{
		ProviderName: "providername",
		AccessToken:  "accesstoken",
		RefreshToken: "refreshtoken",
		Name:         "TestCreateProvider",
		Lastname:     "lastname",
		Firstname:    "firstname",
		UserID:       "userid",
		AvatarURL:    "avatarurl",
		Nickname:     "nickname",
		Location:     "location",
	}
	err := testProvider.Create(Ctx)
	if err != nil {
		t.Errorf("creating a new provider failed: %s", err)
	}
	testProvider = &Provider{}
	if testProvider.Name != "" {
		t.Errorf("testing variable is not empty: %s", testProvider.Name)
	}
	database.DB.Where("name=?", "TestCreateProvider").First(&testProvider)
	if testProvider.Name != "TestCreateProvider" || testProvider.ProviderName != "providername" {
		t.Error("provider was not successfully inserted")
	}
}

func TestFindUserProviderByName(t *testing.T) {
	setupDatabaseTests()
	testProvider := Provider{
		ProviderName: "providername",
		AccessToken:  "accesstoken",
		RefreshToken: "refreshtoken",
		Name:         "TestFindUserProviderByName",
		Lastname:     "lastname",
		Firstname:    "firstname",
		UserID:       "userid",
		AvatarURL:    "avatarurl",
		Nickname:     "nickname",
		Location:     "location",
	}
	testUser := User{
		Name:  "TestFindUserProviderByName",
		UUID:  uuid.New().String(),
		Email: "email@finduser.de",
	}
	database.DB.Create(&testUser)
	testUser.AppendProviderToUser(Ctx, &testProvider)
	returnedProvider, err := FindUserProviderByName(Ctx, testUser, "providername")
	if err != nil {
		t.Errorf("finding provider by username ended with an error: %s", err)
	}
	if returnedProvider.Name != "TestFindUserProviderByName" {
		t.Error("no or wrong provider was found")
	}
}
