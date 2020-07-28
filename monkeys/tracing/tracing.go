package tracing

import (
	"context"
	"fmt"
	"io"
	"regexp"
	"runtime"

	"github.com/labstack/echo"
	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

var (
	// Tracer defines current Tracer instants for jaeger-tracings
	Tracer opentracing.Tracer
	// Closer defines current io-closer for Tracer
	Closer io.Closer
	// REstripFnPreamble defines Regex to extract just the function name (not the module path)
	REstripFnPreamble = regexp.MustCompile(`^.*\.(.*)$`)
)

// InitJaeger returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
// Use it in this way:
// 	// tracer init
//  closer, span, ctx := tracing.InitJaeger("bongo-auth")
//  defer closer.Close()
//  defer span.Finish()
func InitJaeger(service string, params ...string) (io.Closer, opentracing.Span, context.Context) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: true,
		},
	}
	var err error
	Tracer, Closer, err = cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(Tracer)
	var spanName string
	if len(params) > 0 {
		spanName = params[0]
	} else {
		spanName = getFuncName()
	}
	span := Tracer.StartSpan(spanName)
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	return Closer, span, ctx
}

// Enter has to be called at the beginning of any function with available echo.Context to get a new span named by the calling function.
// Usage at the beginning of a func:
// 	 span := tracing.Enter(c)
//   defer span.Finish()
func Enter(c echo.Context) opentracing.Span {
	funcName := getFuncName()
	span, _ := opentracing.StartSpanFromContext(GetContext(c), funcName)
	span.SetTag("func", funcName)
	return span
}

// EnterWithContext has to be called at the beginning of any function with available context.Context to get a new span named by the calling function.
// Usage at the beginning of a func:
// 	 span := tracing.EnterWithContext(ctx)
//   defer span.Finish()
func EnterWithContext(ctx context.Context) opentracing.Span {
	funcName := getFuncName()
	span, _ := opentracing.StartSpanFromContext(ctx, getFuncName())
	span.SetTag("func", funcName)
	return span
}

// GetContext extracts from echo.Context the context.Context
func GetContext(c echo.Context) context.Context {
	return c.Get("tracingctx").(context.Context)
}

// getFuncName returns the name of the current calling function
func getFuncName() string {
	fnName := "<unkown>"
	pc, _, _, ok := runtime.Caller(2)
	if ok {
		fnName = REstripFnPreamble.ReplaceAllString(runtime.FuncForPC(pc).Name(), "$1")
	}
	return fnName
}
