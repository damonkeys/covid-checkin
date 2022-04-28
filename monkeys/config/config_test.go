package config

import (
	"context"
	"os"
	"testing"

	"github.com/damonkeys/covid-checkin/monkeys/tracing"
	"github.com/labstack/echo/v4"
)

type (
	// ServerConfigStruct holds the server-config
	TestStruct struct {
		Test1         string `env:"TEST_ENV1"`
		Test2         string `env:"TEST_ENV2"`
		Test3         string `env:"TEST_ENV3"`
		BoolTestTrue  bool   `env:"BOOLTRUE"`
		BoolTestFalse bool   `env:"BOOLFALSE"`
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
	os.Unsetenv("BOOLTRUE")
	os.Unsetenv("BOOLFALSE")

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
	os.Setenv("BOOLTRUE", "true")
	os.Setenv("BOOLFALSE", "false")
	testInterface, err := ReadEnvVars(Ctx, TestStruct{})
	if err != nil {
		t.Errorf("all environment variables are set but don't validate: %s", err)
	}
	testStruct := testInterface.(TestStruct)
	if testStruct.Test1 != os.Getenv("TEST_ENV1") ||
		testStruct.Test2 != os.Getenv("TEST_ENV2") ||
		testStruct.Test3 != os.Getenv("TEST_ENV3") ||
		testStruct.BoolTestFalse != false ||
		testStruct.BoolTestTrue != true {
		t.Errorf("environment variables not copied correctly to struct")
	}
}

func TestBoolEnvVars(t *testing.T) {
	// we need to set all envvars for valid struct
	os.Setenv("TEST_ENV1", "not important for this test")
	os.Setenv("TEST_ENV2", "not important for this test")
	os.Setenv("TEST_ENV3", "not important for this test")

	os.Setenv("BOOLTRUE", "true")
	os.Setenv("BOOLFALSE", "false")

	// Test true-values
	trueValues := [5]string{"true", "TRUE", "True", "tRue", "TRue"}
	for i := 0; i < len(trueValues); i++ {
		os.Setenv("BOOLTRUE", trueValues[i])

		testInterface, err := ReadEnvVars(Ctx, TestStruct{})
		if err != nil {
			t.Errorf("all environment variables are set but don't validate: %s", err)
		}
		testStruct := testInterface.(TestStruct)
		if testStruct.BoolTestTrue != true {
			t.Errorf("Bool-Var expected to be true but it is false.")
		}
	}

	// Test false-values
	falseValues := [6]string{"false", "FALSE", "False", "fAlse", "faLSe", "anything"}
	for i := 0; i < len(falseValues); i++ {
		os.Setenv("BOOLFALSE", falseValues[i])

		testInterface, err := ReadEnvVars(Ctx, TestStruct{})
		if err != nil {
			t.Errorf("all environment variables are set but don't validate: %s", err)
		}
		testStruct := testInterface.(TestStruct)
		if testStruct.BoolTestFalse != false {
			t.Errorf("Bool-Var expected to be false but it is true.")
		}
	}
}
