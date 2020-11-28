// Package logger provide safe logger wrapper
package logger

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type Level uint32

const (
	emptyLevel Level = iota
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	// config: panic
	PanicLevel
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	// config: fatal
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	// config: error
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	// config: warning
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	// config: info
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	// config: debug
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	// config: trace
	TraceLevel
)

var levelText = map[Level]string{
	PanicLevel: "panic",
	FatalLevel: "fatal",
	ErrorLevel: "error",
	WarnLevel:  "warning",
	InfoLevel:  "info",
	DebugLevel: "debug",
	TraceLevel: "trace",
}

var levelNumber = map[string]Level{
	"panic":   PanicLevel,
	"fatal":   FatalLevel,
	"error":   ErrorLevel,
	"warning": WarnLevel,
	"info":    InfoLevel,
	"debug":   DebugLevel,
	"trace":   TraceLevel,
}

func (level Level) MarshalText() ([]byte, error) {
	if text, ok := levelText[level]; ok {
		return []byte(text), nil
	}
	return nil, fmt.Errorf("not a valid level %d", level)
}

// Convert the Level to a string. E.g. PanicLevel becomes "panic".
func (level Level) String() string {
	if b, err := level.MarshalText(); err == nil {
		return string(b)
	} else {
		return "unknown"
	}
}

func (level *Level) UnmarshalText(text []byte) error {
	l, ok := levelNumber[string(text)]
	if !ok {
		return fmt.Errorf("not a valid level %d", level)
	}

	*level = l

	return nil
}

func (level *Level) UnmarshalYAML(value *yaml.Node) error {
	return level.UnmarshalText([]byte(value.Value))
}

func (level Level) MarshalYAML() (interface{}, error) {
	data, err := level.MarshalText()
	return string(data), err
}
func (level *Level) UnmarshalJSON(value []byte) error {
	var temp string
	err := json.Unmarshal(value, &temp)
	if err != nil {
		return err
	}
	return level.UnmarshalText([]byte(temp))
}

func (level Level) MarshalJSON() ([]byte, error) {
	data, ok := levelText[level]
	if !ok {
		return nil, fmt.Errorf("unlnown level")
	}
	return json.Marshal(data)
}
