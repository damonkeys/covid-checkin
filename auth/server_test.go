package main

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"

	"../monkeys/tracing"
)

var (
	Ctx context.Context
	C   echo.Context
)

func TestParseCommandLineParameter(t *testing.T) {
	initTest()
	e := echo.New()
	readEnvironmentConfig(Ctx, e.Logger)
	if serverConfig.Port != os.Getenv("SERVER_PORT") {
		t.Errorf("wrong port read from default parameter: %s", serverConfig.Port)
	}
}

func TestLogin(t *testing.T) {
	c, _ := setupTest()
	setupDatabaseTests(tracing.GetContext(c))

	err := login(c)
	if err != nil {
		t.Errorf("login failed: %s", err)
	}
}

func TestGetUsernameInvalidSession(t *testing.T) {
	c, rec := setupTest()
	setupDatabaseTests(tracing.GetContext(c))

	err := getLoginStatus(c)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	var responseFailed map[string]interface{}
	json.Unmarshal([]byte(rec.Body.String()), &responseFailed)
	if responseFailed["useronline"] == true {
		t.Error("JSON-response ended successfully and got a username without a session")
	}
	if responseFailed["username"] != "" {
		t.Errorf("JSON-response sends username! %s", responseFailed["username"])
	}
}

func TestGetUsernameValidSession(t *testing.T) {
	// create user and new session
	c, rec := setupTest()
	setupDatabaseTests(tracing.GetContext(c))

	var responseSuccess map[string]interface{}
	gothUser := createGothTestUser("TestCreateSessionCookie", "provider")
	createNewSessionCookie(c, gothUser)
	err := getLoginStatus(c)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	json.Unmarshal([]byte(rec.Body.String()), &responseSuccess)
	if responseSuccess["useronline"] == false {
		t.Error("JSON-response failed with no session and session-user")
	}
	if responseSuccess["username"] == "" {
		t.Error("JSON-response sends NO username!")
	}
}

func TestGetCallbackURL(t *testing.T) {
	initTest()
	baseURL := serverConfig.Baseurl
	resultURL := getCallbackURL(C, "")
	if resultURL != baseURL {
		t.Errorf("callback-URL is %s not %s", resultURL, baseURL)
	}

	resultURL = getCallbackURL(C, "/use")
	if resultURL != baseURL {
		t.Errorf("callback-URL is %s not %s", resultURL, baseURL+"/use")
	}

	resultURL = getCallbackURL(C, "pay")
	if resultURL != baseURL+"/pay" {
		t.Errorf("callback-URL is %s not %s", resultURL, baseURL+"/pay")
	}
}

func TestGetCallbackURLFromSession(t *testing.T) {
	c, _ := setupTest()
	callbackURL := getCallbackURLFromSession(c)
	baseURL := serverConfig.Baseurl
	// test without cookie set
	if callbackURL != baseURL {
		t.Errorf("callback-URL is %s not %s", callbackURL, baseURL)
	}

	// set cookie and test
	sess, _ := session.Get(sessionName, c)
	sess.Options = &sessions.Options{
		Path:     "/auth",
		MaxAge:   30,
		HttpOnly: true,
	}
	sessions.NewCookie("callbackURL", "/use", sess.Options)
	sess.Values["callbackURL"] = "/use"
	sess.Save(c.Request(), c.Response())
	callbackURL = getCallbackURLFromSession(c)
	if callbackURL != baseURL+"/use" {
		t.Errorf("callback-URL is %s not %s", callbackURL, baseURL)
	}
}

func initTest() {
	os.Setenv("SERVER_PORT", "2000")
	os.Setenv("DB_NAME", "test_monkey_auth")
	os.Setenv("DB_HOST", "")
	os.Setenv("DB_USER", "test_auth_user")
	os.Setenv("DB_PASSWORD", "")
	os.Setenv("P_FACEBOOK_KEY", "123456789012345")
	os.Setenv("P_FACEBOOK_SECRET", "aaabbbccc1234567890123")
	os.Setenv("P_GPLUS_KEY", "123456789012-1234567890abcdef1234567890.apps.googleusercontent.com")
	os.Setenv("P_GPLUS_SECRET", "secretsecretsecretsecretsecretsecret")
	os.Setenv("BASE_URL", "https://example.com")
	os.Setenv("SESSION_SECRET", "secretsecretsecretsecretsecretsecret")
	os.Setenv("ACTIVATION_URL", "https://example.com")
	os.Setenv("ACTIVATION_STATE_URL", "https://example.com")

	// tracer init
	_, _, Ctx = tracing.InitMockJaeger("bongo-auth-test")
	e := echo.New()
	C = e.AcquireContext()
	C.Set("tracingctx", Ctx)
}
