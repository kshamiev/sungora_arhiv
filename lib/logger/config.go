package logger

import (
	"sungora/lib/logger/graylog"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Output    string       `yaml:"output" json:"output"`       // enum (stdout | filePathRelative)
	Formatter string       `yaml:"formatter" json:"formatter"` // enum (json|text)
	Level     logrus.Level `yaml:"level" json:"level"`         // enum (error|warning|info|debug|trace)
	IsCaller  bool         `yaml:"is_caller" json:"is_caller"` // bool
	Hooks     Hooks        `yaml:"hooks" json:"hooks"`
}

type Hooks struct {
	Graylog graylog.Config `yaml:"graylog" json:"graylog"`
}

type CtxKey string

const (
	outStdout     = "stdout"
	outStderr     = "stderr"
	outEmpty      = "empty"
	formatterJSON = "json"
	TraceID       = "trace-id"

	CtxTraceID CtxKey = "trace-id"
)
