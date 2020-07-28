package models

import (
	"fmt"
	"testing"

	"../../monkeys/database"
	"github.com/google/uuid"
)

func TestFindUserByEmail(t *testing.T) {
	setupDatabaseTests()

	testEmail := "test@email.org"
	// Create Test User
	testUser := User{
		Email: testEmail,
		Name:  "TestFindUserByEmail",
		UUID:  uuid.New().String(),
	}
	database.DB.Create(&testUser)

	// Test with wrong email: no findings
	wrongEmail := "wrong" + testEmail
	returnedUser, err := FindUserByEmail(Ctx, wrongEmail)
	if err == nil {
		t.Error("error expected but no error was returned")
	}
	if err != nil && err.Error() != fmt.Sprintf("no existing User with email-address %s found", wrongEmail) {
		t.Errorf("searching user by email was ended with unexpected error: %s", err)
	}
	if returnedUser.Name != "" {
		t.Errorf("a user was found but no user was expected: %s", returnedUser.Name)
	}
	// Test with right email
	returnedUser, err = FindUserByEmail(Ctx, testEmail)
	if err != nil {
		t.Errorf("searching user by email was ended with unexpected error: %s", err)
	}
	if returnedUser.Name != "TestFindUserByEmail" {
		t.Errorf("wrong user was found: %s", returnedUser.Name)
	}
	returnedUser, err = FindUserByEmail(Ctx, "")
	if returnedUser.Name != "" {
		t.Errorf("user found with but no email address was commited to search")
	}
}

func TestFindUserByUUID(t *testing.T) {
	setupDatabaseTests()
	uuid := uuid.New().String()

	testUser := User{
		Email: "test@email.org",
		Name:  "TestFindUserByUUID",
		UUID:  uuid,
	}
	database.DB.Create(&testUser)

	// Test with wrong uuid: no findings
	returnedUser, err := FindUserByUUID(Ctx, "wrong uuid")
	if err == nil {
		t.Error("error expected but no error was returned")
	}
	if err != nil && err.Error() != fmt.Sprintf("no existing User with uuid %s found", "wrong uuid") {
		t.Errorf("searching user by uuid was ended with unexpected error: %s", err)
	}
	if returnedUser.Name != "" {
		t.Errorf("no user returning expected but we found an user: %s", returnedUser.Name)
	}
	// Test with valid uuid
	returnedUser, err = FindUserByUUID(Ctx, uuid)
	if err != nil {
		t.Errorf("searching user by uuid was ended with unexpected error: %s", err)
	}
	if returnedUser.Name != "TestFindUserByUUID" {
		t.Errorf("wrong user was found: %s", returnedUser.Name)
	}
}

func TestCreateAndUpdateUser(t *testing.T) {
	setupDatabaseTests()
	testUser := User{
		Email:     "test@email.org",
		Name:      "TestCreateUser",
		UUID:      uuid.New().String(),
		AvatarURL: "https://checkin.chckr.de/static/media/monkeycash-logo.3d72804d.svg",
	}
	// Create
	err := testUser.Create(Ctx)
	if err != nil {
		t.Errorf("creating a new user failed: %s", err)
	}
	testUser = User{}
	if testUser.Name != "" {
		t.Errorf("testing variable is not empty: %s", testUser.Name)
	}
	database.DB.Where("name=?", "TestCreateUser").First(&testUser)
	if testUser.Name != "TestCreateUser" || testUser.Email != "test@email.org" {
		t.Error("user was not successfully inserted")
	}
	// Update
	testUser.Email = "updated@email.org"
	testUser.AvatarURL = "https://update.io"
	testUser.UUID = "CHANGED!"
	err = testUser.Update(Ctx)
	if err == nil {
		t.Error("UUID was changed but user updated without error")
	}
	database.DB.Where("name=?", "TestCreateUser").First(&testUser)
	testUser.Email = "updated@email.org"
	testUser.AvatarURL = "https://update.io"
	err = testUser.Update(Ctx)
	if err != nil {
		t.Errorf("User update failed. Unexpected error: %s", err)
	}
	database.DB.Where("name=?", "TestCreateUser").First(&testUser)
	if testUser.Email != "updated@email.org" {
		t.Errorf("E-Mail not updated. Expected 'updated@email.org', actual: '%s'", testUser.Email)
	}
	if testUser.AvatarURL != "https://update.io" {
		t.Errorf("E-Mail not updated. Expected 'https://update.io', actual: '%s'", testUser.AvatarURL)
	}

}

func TestAppendProviderToUser(t *testing.T) {
	setupDatabaseTests()

	testUser := User{
		Email: "test@email.org",
		Name:  "TestAppendProviderToUser",
		UUID:  uuid.New().String(),
	}
	testProvider := Provider{
		ProviderName: "providername",
		AccessToken:  "accesstoken",
		RefreshToken: "refreshtoken",
		Name:         "TestAppendProviderToUser",
		Lastname:     "lastname",
		Firstname:    "firstname",
		UserID:       "userid",
		AvatarURL:    "avatarurl",
		Nickname:     "nickname",
		Location:     "location",
	}

	// negative test - user was not created
	err := testUser.AppendProviderToUser(Ctx, &testProvider)
	if err == nil {
		t.Error("error expected but no error was returned")
	}

	// positive test
	database.DB.Create(&testUser)
	err = testUser.AppendProviderToUser(Ctx, &testProvider)
	if err != nil {
		t.Errorf("unexpected error was returned: %s", err)
	}

	testUser = User{}
	testProvider = Provider{}

	database.DB.Where("name=?", "TestAppendProviderToUser").First(&testUser)
	if testUser.Name != "TestAppendProviderToUser" {
		t.Error("created testuser not found")
	}
	testProviders := []Provider{}
	database.DB.Model(&testUser).Association("Providers").Find(&testProviders)
	if testProviders[0].Name != "TestAppendProviderToUser" {
		t.Error("created and associated provider not found")
	}
}

func TestDeleteUndeleteUser(t *testing.T) {
	setupDatabaseTests()
	testUser := User{
		Email: "test@email.org",
		Name:  "TestAppendProviderToUser",
		UUID:  uuid.New().String(),
	}
	testProvider := Provider{
		ProviderName: "providername",
		AccessToken:  "accesstoken",
		RefreshToken: "refreshtoken",
		Name:         "TestAppendProviderToUser",
		Lastname:     "lastname",
		Firstname:    "firstname",
		UserID:       "userid",
		AvatarURL:    "avatarurl",
		Nickname:     "nickname",
		Location:     "location",
	}

	database.DB.Create(&testUser)
	err := testUser.AppendProviderToUser(Ctx, &testProvider)
	email := testUser.Email

	dbUser, err := FindUserByEmail(Ctx, email)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if dbUser.Email != email {
		t.Errorf("read email is not the expected: %s != %s", dbUser.Email, email)
	}
	if dbUser.Providers[0].ProviderName != "providername" {
		t.Errorf("provider was'n found.")
	}
	// Soft-delete user
	err = dbUser.Delete(Ctx)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	// Try to find the deleted user
	dbUser, err = FindUserByEmail(Ctx, email)
	if err == nil {
		t.Error("deleted user was found")
	}
	// Read soft-deleted user
	dbUser, err = FindDeletedUserByEmail(Ctx, email)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if dbUser.Email != email {
		t.Errorf("read email is not the expected: %s != %s", dbUser.Email, email)
	}
	// Undelete it!
	err = dbUser.Undelete(Ctx)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	dbUser, err = FindUserByEmail(Ctx, email)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if dbUser.Email != email {
		t.Errorf("read email is not the expected: %s != %s", dbUser.Email, email)
	}
}
