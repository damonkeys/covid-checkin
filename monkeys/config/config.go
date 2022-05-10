package config

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/damonkeys/covid-checkin/monkeys/tracing"
)

// ReadEnvVars checks setting of all environment variable names of the given struct. If they aren't existing an error will returned.
//
// Use it in this way:
//     configInterface, err := config.ReadEnvVars(ServerConfigStruct{})
//     if err != nil {
// 	     e.Logger.Error(err)
// 	     os.Exit(-1)
//     }
//     serverConfig = configInterface.(ServerConfigStruct)
//
func ReadEnvVars(ctx context.Context, configStruct interface{}) (interface{}, error) {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	var unsettedEnvVars []string
	// iterate through struct and read out envvar tag
	// t := reflect.TypeOf(configStruct)
	copyStruct := reflect.New(reflect.ValueOf(configStruct).Type()).Elem()

	iteratedcopyStruct, unsettedEnvVars := iterateStruct(ctx, configStruct, copyStruct)
	copyStruct.Set(iteratedcopyStruct)

	if len(unsettedEnvVars) > 0 {
		return nil, fmt.Errorf("there are missing environment variables: %s", strings.Join(unsettedEnvVars, ", "))
	}
	return copyStruct.Interface(), nil
}

func iterateStruct(ctx context.Context, structBlock interface{}, copyStructBlock reflect.Value) (reflect.Value, []string) {
	span := tracing.EnterWithContext(ctx)
	defer span.Finish()

	var unsettedEnvVars []string
	t := reflect.TypeOf(structBlock)
	copyStruct := reflect.New(reflect.ValueOf(structBlock).Type()).Elem()

	for i := 0; i < t.NumField(); i++ {
		envvar := t.Field(i).Tag.Get("env")

		// Sub-Struct?
		if copyStruct.Field(i).Kind().String() == "struct" {
			iteratedCopyStruct, iteratedUnsetVars := iterateStruct(ctx, copyStruct.Field(i).Interface(), copyStructBlock.Field(i))
			for _, envvar := range iteratedUnsetVars {
				unsettedEnvVars = append(unsettedEnvVars, envvar)
			}
			copyStruct.Field(i).Set(iteratedCopyStruct)
		}
		if envvar != "" {
			if _, exists := os.LookupEnv(envvar); exists {
				fieldType := copyStruct.Field(i).Type().String()
				switch fieldType {
				case "string":
					copyStruct.Field(i).SetString(os.Getenv(envvar))

				case "bool":
					boolValue := (strings.ToLower(os.Getenv(envvar)) == "true")
					copyStruct.Field(i).SetBool(boolValue)
				}
			} else {
				// remember all missing environment variables
				unsettedEnvVars = append(unsettedEnvVars, envvar)
			}
		}
	}
	return copyStruct, unsettedEnvVars
}
