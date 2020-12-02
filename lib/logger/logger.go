// Package logger provide safe logger wrapper
package logger

import (
	"context"
	"io"
	"log"
)

// Fields is a log fields type
type Fields map[string]interface{}

const (
	// fields
	ErrorField = "error"
	TitleField = "title"

	// outputs
	Stdout = "stdout"
	Stderr = "stderr"
	Vacuum = "vacuum"

	// formatters
	JSONFormatter = "json"
	TextFormatter = "text"
)

type StdLogger interface {
}

type Logger interface {
	// Level calls
	Trace(args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(...interface{})
	Panic(...interface{})
	Audit(objectIdentifier, userIdentifier, action string)
	// Native log with level
	Log(level Level, args ...interface{})

	// Writer
	Writer() *io.PipeWriter
	// Standard logger
	CreateStdLogger() *log.Logger
	// Context
	WithField(key string, value interface{}) Logger
	WithFields(fields Fields) Logger
	WithError(err error) Logger
}

// Config is a configuration for logger
type Config struct {
	Title        string `yaml:"title" json:"title" toml:"title"`                      // no field "title" if value is empty
	Output       string `yaml:"output" json:"output" toml:"output"`                   // enum (stdout|stderr|vacuum|path/to/file)
	AuditOutput  string `yaml:"audit_output" json:"audit_output" toml:"audit_output"` // enum (stdout|stderr|vacuum|path/to/file)
	Formatter    string `yaml:"formatter" json:"formatter" toml:"formatter"`          // enum (json|text)
	ReportCaller bool   `yaml:"reportCaller" json:"reportCaller" toml:"reportCaller"` // bool
	Level        Level  `yaml:"level" json:"level" toml:"level"`                      // enum (panic|fatal|error|warning|info|debug|trace)
	Hooks        Hooks  `yaml:"hooks" json:"hooks" toml:"hooks"`
}

// Hooks is a set of hooks for logger
type Hooks struct {
	Sentry *Sentry `yaml:"sentry" json:"sentry" toml:"sentry"`

	// Syslog   *Syslog   `yaml:"syslog" json:"syslog" toml:"syslog"`
}

type ctxlog struct{}

// WithLogger put logger to context
func WithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, ctxlog{}, logger)
}

var DefaultLogger Logger = newLogrusWrapper(nil)

// GetLogger get logger from context, or DefaultLogger if not exists
func GetLogger(ctx context.Context) Logger {
	return GetLoggerDefault(ctx, DefaultLogger)
}

// GetLoggerDefault get logger from context
// defLogger - logger that is used if logger is not found in context
func GetLoggerDefault(ctx context.Context, defLogger Logger) Logger {
	l, ok := ctx.Value(ctxlog{}).(Logger)
	if !ok {
		l = defLogger
	}
	return l
}
