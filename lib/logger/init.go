package logger

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var instance Logger

// from config
func NewLogger(config *Config) Logger {
	return newLogrusWrapper(config)
}

func Init(config *Config) Logger {
	instance = newLogrusWrapper(config)
	return instance
}

var mu sync.RWMutex

func Get() Logger {
	if instance == nil {
		mu.Lock()
		if instance == nil {
			Init(&Config{
				Title:        "Default",
				Output:       "stdout",
				AuditOutput:  "stdout",
				Formatter:    "json",
				ReportCaller: false,
				Level:        TraceLevel,
				Hooks:        Hooks{},
			})
		}
		mu.Unlock()
	}
	return instance
}

var formatters = map[string]logrus.Formatter{
	JSONFormatter: &logrus.JSONFormatter{},
	TextFormatter: &logrus.TextFormatter{},
}
