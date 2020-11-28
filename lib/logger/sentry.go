package logger

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/raven-go"
)

// Sentry hook configuration
type Sentry struct {
	DSN                               string            `yaml:"dsn" json:"dns"`
	Level                             Level             `yaml:"level" json:"level" toml:"level"`
	Tags                              map[string]string `yaml:"tags" json:"tags"`
	Timeout                           time.Duration     `yaml:"timeout" json:"timeout"`
	StacktraceConfigurationContext    int               `yaml:"stacktrace_context" json:"stacktrace_context"`
	StacktraceConfigurationLevel      Level             `yaml:"stacktrace_level" json:"stacktrace_level"`
	StacktraceConfigurationEnable     bool              `yaml:"stacktrace_enable" json:"stacktrace_enable"`
	StacktraceConfigurationBreadcrumb bool              `yam:"stacktrace_breadcrumb" json:"stacktrace_breadcrumb"`
	Async                             bool              `yaml:"async" json:"async"`
	SSLSkipVerify                     bool              `yaml:"ssl_skip_verify" json:"ssl_skip_verify"`
}

func logLevels(config *Sentry) []logrus.Level {
	// TODO must
	levels := make([]logrus.Level, 0, config.Level+1)
	for i := PanicLevel; i <= config.Level; i++ {
		levels = append(levels, logrusLevelMap[i])
	}
	return levels
}

func sentryHook(config *Sentry) (*logrus_sentry.SentryHook, error) {
	var (
		hook *logrus_sentry.SentryHook
		err  error
	)
	levels := logLevels(config)
	client, err := raven.New(config.DSN)
	if err != nil {
		return nil, err
	}
	if len(config.Tags) != 0 {
		client.Tags = config.Tags
	}

	if config.SSLSkipVerify {
		client.Transport.(*raven.HTTPTransport).Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify = true
	}

	if config.Async {
		hook, err = logrus_sentry.NewAsyncWithClientSentryHook(client, levels)
	} else {
		hook, err = logrus_sentry.NewWithClientSentryHook(client, levels)
	}

	if config.Timeout == 0 {
		hook.Timeout = 2 * time.Second
	} else {
		hook.Timeout = config.Timeout
	}

	hook.StacktraceConfiguration.Enable = config.StacktraceConfigurationEnable
	hook.StacktraceConfiguration.Level = logrusLevelMap[config.StacktraceConfigurationLevel]
	hook.StacktraceConfiguration.Context = config.StacktraceConfigurationContext
	hook.StacktraceConfiguration.IncludeErrorBreadcrumb = config.StacktraceConfigurationBreadcrumb

	return hook, err
}
