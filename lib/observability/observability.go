package observability

import (
	"context"
	"net/http"

	"contrib.go.opencensus.io/exporter/jaeger"
	"contrib.go.opencensus.io/exporter/prometheus"
	"go.opencensus.io/plugin/runmetrics"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"

	"sungora/lib/logger"
)

type Config struct {
	Title   string            `yaml:"title"`
	Tags    map[string]string `yaml:"tags"`
	Tracing struct {
		Jaeger *struct {
			AgentAddress     string `yaml:"agent_address"`
			CollectorAddress string `yaml:"collector_address"`
			User             string `yaml:"user"`
			Password         string `yaml:"password"`
			ServiceName      string `yaml:"service_name"`
		} `yaml:"jaeger"`
		Probability float64 `yaml:"probability"`
	} `yaml:"tracing"`
	Monitoring struct {
		Prometheus *struct {
			Address   string `yaml:"address"`
			Namespace string `yaml:"namespace"`
		} `yaml:"prometheus"`
	} `yaml:"monitoring"`
}

func Enable(ctx context.Context, config *Config) {
	log := logger.GetLogger(ctx).WithFields(map[string]interface{}{
		"operation": "enableObservabilityAndExporters",
	})
	jaegerTags := []jaeger.Tag{}
	for k, v := range config.Tags {
		jaegerTags = append(jaegerTags, jaeger.StringTag(k, v))
	}
	if config.Tracing.Jaeger == nil {
		return
	}
	je, err := jaeger.NewExporter(jaeger.Options{
		CollectorEndpoint: config.Tracing.Jaeger.CollectorAddress,
		Process: jaeger.Process{
			ServiceName: config.Title,
			Tags:        jaegerTags,
		},
		AgentEndpoint: config.Tracing.Jaeger.AgentAddress,
		Username:      config.Tracing.Jaeger.User,
		Password:      config.Tracing.Jaeger.Password,
	})

	if err != nil {
		log.WithError(err).Error("failed to create the Jaeger exporter")
	}
	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(config.Tracing.Probability)})
	if config.Monitoring.Prometheus == nil {
		return
	}
	pe, err := prometheus.NewExporter(prometheus.Options{
		Namespace:   config.Monitoring.Prometheus.Namespace,
		ConstLabels: config.Tags,
	})
	if err != nil {
		log.WithError(err).Error("failed to create the Prometheus stats exporter")
	}
	_ = runmetrics.Enable(runmetrics.RunMetricOptions{
		EnableCPU:    true,
		EnableMemory: true,
	})

	view.RegisterExporter(pe)
	if ochttpViewErr := view.Register(AllHTTPViews...); ochttpViewErr != nil {
		log.WithError(ochttpViewErr).Error("failed to register server views for HTTP metrics")
	}
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", pe)
		if promErr := http.ListenAndServe(config.Monitoring.Prometheus.Address, mux); err != nil {
			log.WithError(promErr).Error("Failed to run Prometheus scrape endpoint")
		}
	}()
}
