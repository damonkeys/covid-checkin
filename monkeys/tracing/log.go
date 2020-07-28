package tracing

import (
	"encoding/json"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

// LogStruct marshalls a struct to JSON and log it as a string to the given span
func LogStruct(span opentracing.Span, key string, dataStruct interface{}) {
	marshaledStruct, err := json.Marshal(dataStruct)
	if err != nil {
		return
	}
	span.LogFields(log.String(key, string(marshaledStruct)))
}

// LogString writes a logstring to the given span
func LogString(span opentracing.Span, key string, value string) {
	span.LogFields(log.String(key, value))
}

// LogError writes an error to the given span
func LogError(span opentracing.Span, err error) {
	span.LogFields(log.Error(err))
	span.SetTag("error", true)
}

// LogErrorMsg writes a message and an error to the given span
func LogErrorMsg(span opentracing.Span, err error, msg string) {
	LogString(span, "message", msg)
	LogError(span, err)
}
