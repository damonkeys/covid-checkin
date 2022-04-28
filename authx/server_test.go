package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/markbates/goth"

	"github.com/gorilla/sessions"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	"github.com/damonkeys/covid-checkin/monkeys/tracing"
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

	err := login(c)

	if err != nil {
		t.Errorf("login failed: %s", err)
	}
}

func TestGetUsernameInvalidSession(t *testing.T) {
	c, rec := setupTest()

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

func TestGetCallbackURL(t *testing.T) {
	initTest()
	baseURL := serverConfig.Baseurl

	resultURL := getCallbackURL(C, "")

	if resultURL != baseURL {
		t.Errorf("callback-URL is %s not %s", resultURL, baseURL)
	}

	resultURL = getCallbackURL(C, "/use")
	if resultURL != baseURL+"/use" {
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

func TestBasiGothInitialisation(t *testing.T) {
	initGoth()

	providers := goth.GetProviders()

	if len(providers) != 3 {
		t.Errorf("expected provider count to be 3 (fb, google, apple) but insted found: %s", providers)
	}
}

func TestIsExpectedPostRequest(t *testing.T) {
	t.Parallel()
	values := url.Values{}

	expectTrue := isExpectedPostRequest(values, "POST")

	if expectTrue == false {
		t.Fatal("empty post request should return true")
	}
}

func TestIsExpectedPostRequestWithGet(t *testing.T) {
	t.Parallel()
	values := url.Values{}

	expectFalse := isExpectedPostRequest(values, "GET")

	if expectFalse == true {
		t.Fatal("empty GET request should return false")
	}
}

func TestIsExpectedPostRequestWithFantasyMethod(t *testing.T) {
	t.Parallel()
	values := url.Values{}

	expectFalse := isExpectedPostRequest(values, "FANTASY")

	if expectFalse == true {
		t.Fatal("empty FANTASY request should return false")
	}
}

func TestIsExpectedPostRequestWithValues(t *testing.T) {
	t.Parallel()
	values := url.Values{"foo": {"bar"}}

	expectFalse := isExpectedPostRequest(values, "FANTASY")

	if expectFalse == true {
		t.Fatal("POST request with values should return false")
	}
}

func TestResolveUserNameFromRequestIfApple(t *testing.T) {
	// idea from here: https: //www.reddit.com/r/golang/comments/6bortg/how_to_test_post_requests_with_data_other_than/
	t.Parallel()
	e := echo.New()
	userJSON := `{"name":{"firstName":"first", "lastName": "last"}, "email":"foo@example.com"}`
	form := url.Values{}
	form.Add("user", userJSON)
	req, err := http.NewRequest("POST", "http://example.com", strings.NewReader(form.Encode()))
	req.Form = form
	if err != nil {
		t.Fatalf("Couldn't create request. Error: %s", err)
	}
	c := e.NewContext(req, nil)
	context := context.TODO()
	c.Set("tracingctx", context)
	gothUser := &goth.User{}
	gothUser.Provider = "apple"

	err = resolveUserNameFromRequestIfApple(c, gothUser)

	if gothUser.FirstName != "first" {
		t.Fatalf("Wrong first name: %s", gothUser.FirstName)
	}
	if gothUser.LastName != "last" {
		t.Fatalf("Wrong first name: %s", gothUser.LastName)
	}
	if gothUser.Name != "first last" {
		t.Fatalf("Wrong name: %s", gothUser.Name)
	}
}

func TestResolveWrongUserNameFromRequestIfApple(t *testing.T) {
	// idea from here: https: //www.reddit.com/r/golang/comments/6bortg/how_to_test_post_requests_with_data_other_than/
	t.Parallel()
	e := echo.New()
	userJSON := `{"this_is_invalid"}`
	form := url.Values{}
	form.Add("user", userJSON)
	req, err := http.NewRequest("POST", "http://example.com", strings.NewReader(form.Encode()))
	req.Form = form
	if err != nil {
		t.Fatalf("Couldn't create request. Error: %s", err)
	}
	c := e.NewContext(req, nil)
	context := context.TODO()
	c.Set("tracingctx", context)
	gothUser := &goth.User{}
	gothUser.Provider = "apple"

	err = resolveUserNameFromRequestIfApple(c, gothUser)

	if err == nil {
		t.Fatalf("expected an  error on invalid json. Error: %s", err)
	}
}

func TestResolveAppleUserNamesOnlyOnAppleProvidedUser(t *testing.T) {
	// idea from here: https: //www.reddit.com/r/golang/comments/6bortg/how_to_test_post_requests_with_data_other_than/
	t.Parallel()
	e := echo.New()
	userJSON := `{"name":{"firstName":"first", "lastName": "last"}, "email":"foo@example.com"}`
	form := url.Values{}
	form.Add("user", userJSON)
	req, err := http.NewRequest("POST", "http://example.com", strings.NewReader(form.Encode()))
	req.Form = form
	if err != nil {
		t.Fatalf("Couldn't create request. Error: %s", err)
	}
	c := e.NewContext(req, nil)
	context := context.TODO()
	c.Set("tracingctx", context)
	gothUser := &goth.User{}
	gothUser.Provider = "xxx"

	err = resolveUserNameFromRequestIfApple(c, gothUser)

	if err != nil {
		t.Fatalf("did not expect an error on valid json with xxx provider. Error: %s", err)
	}
	if gothUser.FirstName != "" {
		t.Fatalf("gothUser cannot have a FirstName attribute (from apple) as this should not be parsed"+
			" on a non apple provided request: %s", gothUser)
	}
	if gothUser.LastName != "" {
		t.Fatalf("gothUser cannot have a LastName attribute (from apple) as this should not be parsed"+
			" on a non apple provided request: %s", gothUser)
	}
	if gothUser.Name != "" {
		t.Fatalf("gothUser cannot have a Name attribute (from apple) as this should not be created"+
			" on a non apple provided request: %s", gothUser)
	}
}

func initTest() {
	os.Setenv("SERVER_PORT", "2000")
	os.Setenv("DB_NAME", "test_monkey_auth")
	os.Setenv("DB_HOST", "")
	os.Setenv("DB_USER", "test_auth_user")
	os.Setenv("DB_PASSWORD", "monkey_auth_pw")
	os.Setenv("P_FACEBOOK_KEY", "123456789012345")
	os.Setenv("P_FACEBOOK_SECRET", "aaabbbccc1234567890123")
	os.Setenv("P_GPLUS_KEY", "123456789012-1234567890abcdef1234567890.apps.googleusercontent.com")
	os.Setenv("P_GPLUS_SECRET", "secretsecretsecretsecretsecretsecret")
	os.Setenv("P_APPLE_KEY", "de.chckr.checkin")
	os.Setenv("P_APPLE_SECRET", "secretsecretsecretsecretsecretsecret")
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
