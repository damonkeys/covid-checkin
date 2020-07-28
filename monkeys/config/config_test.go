package config

import (
	"context"
	"os"
	"testing"

	"../tracing"
	"github.com/labstack/echo"
)

type (
	// ServerConfigStruct holds the server-config
	TestStruct struct {
		Test1 string `env:"TEST_ENV1"`
		Test2 string `env:"TEST_ENV2"`
		Test3 string `env:"TEST_ENV3"`
	}
)

var (
	Ctx context.Context
	C   echo.Context
)

func init() {
	// tracer init
	_, _, Ctx = tracing.InitMockJaeger("bongo-auth-test")
	e := echo.New()
	C = e.AcquireContext()
	C.Set("tracingctx", Ctx)
}

func TestValidateEnvVars(t *testing.T) {
	// negative test - no setted environment variables!
	os.Unsetenv("TEST_ENV1")
	os.Unsetenv("TEST_ENV2")
	os.Unsetenv("TEST_ENV3")

	_, err := ReadEnvVars(Ctx, TestStruct{})
	if err == nil {
		t.Error("all environment variables seemed to be set but are deleted. something went wrong")
	}

	// set only one variable
	os.Setenv("TEST_ENV2", "test content")
	_, err = ReadEnvVars(Ctx, TestStruct{})
	if err == nil {
		t.Error("all environment variables seemed to be set but only one was set. something went wrong")
	}

	// set all variables
	os.Setenv("TEST_ENV1", "more test content")
	os.Setenv("TEST_ENV3", "last test content")
	testInterface, err := ReadEnvVars(Ctx, TestStruct{})
	if err != nil {
		t.Errorf("all environment variables are set but don't validate: %s", err)
	}
	testStruct := testInterface.(TestStruct)
	if testStruct.Test1 != os.Getenv("TEST_ENV1") ||
		testStruct.Test2 != os.Getenv("TEST_ENV2") ||
		testStruct.Test3 != os.Getenv("TEST_ENV3") {
		t.Errorf("environment variables not copied correctly to struct")
	}
}
