package logger

import (
	"github.com/sirupsen/logrus"
)

// CreateLogger from config
func CreateLogger(config *Config) Logger {
	logger := newLogrusWrapper(config)
	return logger
}

var formatters = map[string]logrus.Formatter{
	JSONFormatter: &logrus.JSONFormatter{},
	TextFormatter: &logrus.TextFormatter{},
}
