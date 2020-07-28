package tracing

import (
	"context"
	"fmt"
	"io"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/mocktracer"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

// InitMockJaeger returns an instance of mocked Jaeger Tracer.
func InitMockJaeger(service string, params ...string) (io.Closer, opentracing.Span, context.Context) {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans: false,
		},
	}
	var err error
	Tracer := mocktracer.New()
	_, Closer, err = cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Mock-Jaeger: %v\n", err))
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
