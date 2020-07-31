package main

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/damonkeys/ch3ck1n/auth/models"
	"github.com/damonkeys/ch3ck1n/monkeys/database"
	"github.com/damonkeys/ch3ck1n/monkeys/tracing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	log "github.com/labstack/gommon/log"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func TestCreateSessionCookie(t *testing.T) {
	// Setup
	c, _ := setupTest()
	setupDatabaseTests(tracing.GetContext(c))

	// session cookie availabe before
	sess, _ := session.Get("_monkeycash_session", c)
	if sess.Values["userid"] != nil {
		t.Errorf("userid %s read from session cookie. it might be empty", sess.Values["userid"])
	}
	gothUser := createGothTestUser("TestCreateSessionCookie", "provider")
	err := createNewSessionCookie(c, gothUser)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if sess.Values["userid"] == nil {
		t.Error("userid was not set in session-cookie")
	}
	// new provider
	gothUser = createGothTestUser("TestCreateSessionCookie", "newProvider")
	createNewSessionCookie(c, gothUser)
	dbUser, _ := models.FindUserByEmail(tracing.GetContext(c), gothUser.Email)
	dbProvider, err := models.FindUserProviderByName(tracing.GetContext(c), *dbUser, "newProvider")
	if err != nil {
		t.Errorf("unexpected error: %s\n", err)
	}
	if dbProvider == nil {
		t.Error("provider not found")
	}
	// test with soft-deleted user
	uuid := dbUser.UUID
	err = dbUser.Delete(tracing.GetContext(c))
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	dbUser, err = models.FindUserByEmail(tracing.GetContext(c), gothUser.Email)
	if err == nil {
		t.Error("deleted user was found")
	}
	if dbUser.Email == gothUser.Email {
		t.Error("user was not soft-deleted!")
	}
	err = createNewSessionCookie(c, gothUser)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	dbUser, err = models.FindUserByEmail(tracing.GetContext(c), gothUser.Email)
	if dbUser.UUID != uuid {
		t.Errorf("soft-deleted user not undeleted. new user was created!")
	}
}

func TestRemoveSessionCookie(t *testing.T) {
	c, _ := setupTest()
	setupDatabaseTests(tracing.GetContext(c))
	createNewSessionCookie(c, createGothTestUser("TestRemoveSessionCookie", "povider"))
	sess, _ := session.Get("_monkeycash_session", c)
	if sess.Values["userid"] == nil {
		t.Error("no userid read from session cookie")
	}
	// testing function
	err := removeSessionCookie(c)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestFindSessionUser(t *testing.T) {
	// Setup
	c, _ := setupTest()
	setupDatabaseTests(tracing.GetContext(c))

	// No valid session
	dbUser, err := findSessionUser(c)
	if err.Error() != "no valid session found" {
		t.Errorf("unexpected error: %s", err)
	}
	if dbUser.Name != "" {
		t.Error("user found in empty database")
	}

	// create session
	gothUser := createGothTestUser("TestFindSessionUser", "provider")
	err = createNewSessionCookie(c, gothUser)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	// test again
	dbUser, err = findSessionUser(c)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if dbUser.Name != "TestFindSessionUser" {
		t.Errorf("wrong user found: %s", dbUser.Name)
	}
	// test with changed and unkown uuid
	sess, _ := session.Get("_monkeycash_session", c)
	sess.Values["userid"] = "manipulated"
	dbUser, err = findSessionUser(c)
	if err.Error() != fmt.Sprintf("User with userid %s not found - no existing User with uuid %s found", "manipulated", "manipulated") {
		t.Errorf("unexpected error: %s", err)
	}
	if dbUser.Name != "" {
		t.Error("user with wrong uuid found in database")
	}
}
func TestCreateNewProviderData(t *testing.T) {
	c, _ := setupTest()
	gothUser := createGothTestUser("TestCreateNewProviderData", "provider")
	dbProvider := createNewProviderData(c, gothUser)
	if dbProvider.ProviderName != gothUser.Provider {
		t.Error("ProviderName was not set successfully")
	}
	if dbProvider.AccessToken != gothUser.AccessToken {
		t.Error("AccessToken was not set successfully")
	}
	if dbProvider.RefreshToken != gothUser.RefreshToken {
		t.Error("RefreshToken was not set successfully")
	}
	if dbProvider.Name != gothUser.Name {
		t.Error("Name was not set successfully")
	}
	if dbProvider.Lastname != gothUser.LastName {
		t.Error("Lastname was not set successfully")
	}
	if dbProvider.Firstname != gothUser.FirstName {
		t.Error("Firstname was not set successfully")
	}
	if dbProvider.UserID != gothUser.UserID {
		t.Error("UserID was not set successfully")
	}
	if dbProvider.AvatarURL != gothUser.AvatarURL {
		t.Error("AvatarURL was not set successfully")
	}
	if dbProvider.Nickname != gothUser.NickName {
		t.Error("Nickname was not set successfully")
	}
	if dbProvider.Location != gothUser.Location {
		t.Error("Location was not set successfully")
	}
	if dbProvider.ExpiresAt != gothUser.ExpiresAt {
		t.Error("ExpiresAt was not set successfully")
	}
}

func createGothTestUser(funcName string, providerName string) goth.User {
	return goth.User{
		Email:        "test@session.com",
		Provider:     providerName,
		AccessToken:  "accessToken",
		RefreshToken: "refreshToken",
		Name:         funcName,
		LastName:     "lastname",
		FirstName:    "firstname",
		UserID:       "userid",
		AvatarURL:    "avatarurl",
		NickName:     "nickname",
		Location:     "location",
		Description:  "description",
		ExpiresAt:    time.Now(),
	}
}

func setupTest() (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.OFF)
	sessionStore := sessions.NewCookieStore([]byte("secret"))
	gothic.Store = sessionStore
	e.Use(session.Middleware(sessionStore))
	rec := httptest.NewRecorder()
	c := e.NewContext(httptest.NewRequest(echo.GET, "/", nil), rec)
	c.Set("_session_store", sessionStore)

	// tracer init
	_, _, ctx := tracing.InitMockJaeger("bongo-auth-test")
	c.Set("tracingctx", ctx)

	return c, rec
}

// Connects to DBMS, drops alle test-tables and creates them agagin
func setupDatabaseTests(ctx context.Context) {
	databaseConfig := database.ConfigStruct{
		Name:     "test_monkey_auth",
		Password: "",
		User:     "auth_user",
		Server:   "",
	}
	database.InitDatabase(databaseConfig)
	database.DB.DropTableIfExists(&models.User{}, &models.Provider{})
	database.InitDatabase(databaseConfig)
}
