package logger

import (
	"sample/lib/logger/graylog"
)

type Config struct {
	Title     string `yaml:"title" json:"title"`         // title
	Output    string `yaml:"output" json:"output"`       // enum (stdout | filePathRelative)
	Formatter string `yaml:"formatter" json:"formatter"` // enum (json|text)
	Level     string `yaml:"level" json:"level"`         // enum (error|warning|info|debug|trace)
	IsCaller  bool   `yaml:"is_caller" json:"is_caller"` // bool
	Hooks     Hooks  `yaml:"hooks" json:"hooks"`
}

type Hooks struct {
	Graylog graylog.Config `yaml:"graylog" json:"graylog"`
}

type CtxKey string

const (
	OutStdout     = "stdout"
	OutStderr     = "stderr"
	OutEmpty      = "empty"
	FormatterJSON = "json"
	TraceID       = "trace-id"
	Title         = "title"
	Api           = "api"

	CtxTraceID CtxKey = "trace-id"
)
