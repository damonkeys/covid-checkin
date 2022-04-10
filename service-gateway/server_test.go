package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestParseProxyConfig(t *testing.T) {
	os.Setenv("SERVER_PORT_SSL", "443")
	os.Setenv("ROUTES_CONFIG", "./routes.json")
	// test default flags
	setProxyConfig(context.TODO())
	if proxyConfig.ServerPort != "443" {
		t.Errorf("Default port 443 is not set! Actually default port: %s", proxyConfig.ServerPort)
	}
	if proxyConfig.RoutesConfigFile != "./routes.json" {
		t.Errorf("Default route-config filepath is not set! Acually deault filepath: %s", proxyConfig.RoutesConfigFile)
	}
}

func TestParseJSONConfig(t *testing.T) {
	jsonTestfile := "./test-fixtures/routesTest.json"
	proxyConfig.RoutesConfigFile = jsonTestfile
	e := echo.New()

	testParsedRoutes(e, t)
}

// A test for not completed json-files but they are valid
func TestParseJSONConfigNotCompleted(t *testing.T) {
	jsonTestfile := "./test-fixtures/routesTestNotCompleted.json"
	proxyConfig.RoutesConfigFile = jsonTestfile
	e := echo.New()

	testParsedRoutes(e, t)
}

func TestParseNotValidJSON(t *testing.T) {
	jsonTestfile := "./test-fixtures/routesTestNotCompletedNotValidJSON.json"
	proxyConfig.RoutesConfigFile = jsonTestfile
	_, err := ioutil.ReadFile(proxyConfig.RoutesConfigFile)
	if err == nil {
		t.Fatalf("Parsing of not valid JSON test-routes-config successed! What's wrong?")
	}
	t.Log(err)
}

func TestParseIncorrectJSON(t *testing.T) {
	jsonTestfile := "./test-fixtures/routesTestNotValid.json"
	proxyConfig.RoutesConfigFile = jsonTestfile

	var testRoutes *RoutesStruct
	testData, err := ioutil.ReadFile(proxyConfig.RoutesConfigFile)
	if err != nil {
		t.Fatalf("Cannot parse test-routes-config. Error: %s", err)
	}
	json.Unmarshal(testData, &testRoutes)
	valid, errs := validateJSON(testRoutes)
	if valid {
		t.Error("Parsing of different JSON-file with no errors. There must be an error!")
	}
	if len(errs) == 0 {
		t.Error("Parsing of different JSON-file with no validation-errors. There must be an error")
	}
	for _, errElement := range errs {
		t.Logf("Found error: %s", errElement)
	}
}

func TestSimpleUserNamePasswordCheck(t *testing.T) {
	routeStruct := RouteStruct{}
	echoContext := echo.New().NewContext(nil, nil)
	echoContext.Set("tracingctx", context.Background())
	expectFalseWhenRouteEmpty, _ := routeStruct.simpleUserNamePasswordCheck("foo", "bar", echoContext)
	if expectFalseWhenRouteEmpty {
		t.Fatalf("Expect false as return value when using empty route struct but got: %t", expectFalseWhenRouteEmpty)
	}

	routeStruct.Auth.User = "admin"
	routeStruct.Auth.Password = "password"

	expectFalseWhenWrongCredentials, _ := routeStruct.simpleUserNamePasswordCheck("foo", "bar", echoContext)

	if expectFalseWhenWrongCredentials {
		t.Fatalf("Expect false as return value when using wrong credentials but got: %t", expectFalseWhenWrongCredentials)
	}

	expectFalseWhenUsingWrongPassword, _ := routeStruct.simpleUserNamePasswordCheck("admin", "bar", echoContext)

	if expectFalseWhenUsingWrongPassword {
		t.Fatalf("Expect false as return value when using wrong password but got: %t", expectFalseWhenUsingWrongPassword)
	}

	expectFalseWhenUsingWrongUsername, _ := routeStruct.simpleUserNamePasswordCheck("foo", "password", echoContext)

	if expectFalseWhenUsingWrongUsername {
		t.Fatalf("Expect false as return value when using wrong username but got: %t", expectFalseWhenUsingWrongUsername)
	}

	expectTrueWithCorrectCredentials, _ := routeStruct.simpleUserNamePasswordCheck("admin", "password", echoContext)

	if !expectTrueWithCorrectCredentials {
		t.Fatalf("Expect true as return value when using correct credentials: %t", expectTrueWithCorrectCredentials)
	}
}

func testParsedRoutes(e *echo.Echo, t *testing.T) {
	var testRoutes *RoutesStruct
	testData, err := ioutil.ReadFile(proxyConfig.RoutesConfigFile)
	if err != nil {
		t.Fatalf("Cannot parse test-routes-config. Error: %s", err)
	}
	err = json.Unmarshal(testData, &testRoutes)
	// test the parse func
	parseJSONConfig(context.TODO(), e)
	// test jsonfile struct against parsed data
	for _, testRoute := range testRoutes.Routes {
		if !checkAvailablePath(testRoute.Path, e) {
			t.Errorf("Proxy-route '%s' not found!", testRoute.Path)
		}
	}
}

func checkAvailablePath(path string, e *echo.Echo) bool {
	for _, element := range e.Routes() {
		if element.Path == path {
			return true
		}
	}
	return false
}

func validateJSON(testRoutes *RoutesStruct) (bool, []error) {
	var jsonErrors []error
	validFile := true
	if testRoutes.Routes == nil {
		validFile = false
		jsonErrors = append(jsonErrors, errors.New("routes-element not available in JSON-file"))
	}
	if testRoutes.Routes[0].Urls == nil {
		validFile = false
		jsonErrors = append(jsonErrors, errors.New("urls-element not available in JSON-file"))
	}
	if len(testRoutes.Routes[0].Urls) == 0 {
		validFile = false
		jsonErrors = append(jsonErrors, errors.New("urls-elementn with size 0 JSON-file"))
	}
	if testRoutes.Routes[0].Path == "" {
		validFile = false
		jsonErrors = append(jsonErrors, errors.New("path-elementnot available in JSON-file"))
	}
	return validFile, jsonErrors
}
