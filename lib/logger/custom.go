package logger

import (
	"context"

	"contrib.go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
)

func Init(config *Config) Logger {
	DefaultLogger = newLogrusWrapper(config)
	return DefaultLogger
}

func Gist(ctx context.Context) Logger {
	return GetLogger(ctx)
}

// ////

type JaegerConfig struct {
	Title    string  `json:"title" yaml:"title"`
	OcUrl    string  `json:"oc_url" yaml:"oc_url"`
	Fraction float64 `json:"fraction" yaml:"fraction"`
}

type Jaeger struct {
	flush func()
}

// NewJaeger ocTracing = "http://localhost:14268/api/traces"
func NewJaeger(cfg *JaegerConfig) (*Jaeger, error) {
	if err := view.Register(ochttp.DefaultServerViews...); err != nil {
		return nil, err
	}
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		return nil, err
	}
	if err := view.Register(ocgrpc.DefaultClientViews...); err != nil {
		return nil, err
	}

	je, err := jaeger.NewExporter(jaeger.Options{
		CollectorEndpoint: cfg.OcUrl,
		Process: jaeger.Process{
			ServiceName: cfg.Title,
			Tags:        []jaeger.Tag{jaeger.StringTag("system", cfg.Title)},
		},
	})
	if err != nil {
		return nil, err
	}
	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(cfg.Fraction)})
	return &Jaeger{flush: func() {
		je.Flush()
	}}, nil
}

func (comp *Jaeger) Close() {
	comp.flush()
}
