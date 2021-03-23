package main

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/damonkeys/ch3ck1n/monkeys/tracing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	log "github.com/labstack/gommon/log"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

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

func TestFetchNameFromGothUserWithNoName(t *testing.T) {
	t.Parallel()
	gothUser := goth.User{}

	expectSpaceAsName := fetchNameFromGothUser(gothUser)

	if expectSpaceAsName != " " {
		t.Fatalf("Name isn't space as expected: %s", expectSpaceAsName)
	}
}

func TestFetchNameFromGothUserWithName(t *testing.T) {
	t.Parallel()
	gothUser := goth.User{
		Name: "This is the name",
	}

	expectValidName := fetchNameFromGothUser(gothUser)

	if expectValidName != "This is the name" {
		t.Fatalf("Name isn't 'This is the name' as expected: %s", expectValidName)
	}
}

func TestFetchNameFromGothUserWithFirstAndLastName(t *testing.T) {
	t.Parallel()
	gothUser := goth.User{
		FirstName: "FirstName",
		LastName:  "LastName",
	}

	expectValidName := fetchNameFromGothUser(gothUser)

	if expectValidName != "FirstName LastName" {
		t.Fatalf("Name isn't 'FirstName LastName' as expected: %s", expectValidName)
	}
}

func TestFetchNameFromGothUserWithNameAndFirstAndLastName(t *testing.T) {
	t.Parallel()
	gothUser := goth.User{
		Name:      "Name comes first",
		FirstName: "FirstName",
		LastName:  "LastName",
	}

	expectValidName := fetchNameFromGothUser(gothUser)

	if expectValidName != "Name comes first" {
		t.Fatalf("Name isn't 'Name comes first' as expected: %s", expectValidName)
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
