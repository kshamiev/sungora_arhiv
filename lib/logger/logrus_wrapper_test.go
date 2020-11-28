package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"gopkg.in/yaml.v3"
)

//
type testCase struct {
	Message  string
	LogLevel Level
}

func generateLogs(t *testing.T) []testCase {
	t.Helper()
	return []testCase{
		{
			"Info text",
			InfoLevel,
		},
		{
			"Warn text",
			WarnLevel,
		},
		{
			"Err text",
			ErrorLevel,
		},
		{
			"Debug text",
			DebugLevel,
		},
	}
}
func TestCreateStdout(t *testing.T) {
	for _, tc := range generateLogs(t) {
		c := Config{
			Title:  "testTitle",
			Output: "stdout",
			Level:  tc.LogLevel,
		}

		testLogger := NewLogger(&c)
		testLogger.Info(tc.LogLevel, "Info text")
		testLogger.Warning(tc.LogLevel, " Warn text")
		testLogger.Error(tc.LogLevel, "Err text")
		testLogger.Debug(tc.LogLevel, "Debug text")
	}
}

func TestCreateNil(t *testing.T) {
	for _, tc := range generateLogs(t) {
		testLogger := NewLogger(nil)
		testLogger.Info(tc.LogLevel, "Info text")
		testLogger.Warning(tc.LogLevel, " Warn text")
		testLogger.Error(tc.LogLevel, "Err text")
		testLogger.Debug(tc.LogLevel, "Debug text")
	}
}
func TestCtx(t *testing.T) {
	tc := generateLogs(t)[0]
	c := Config{
		Title:  "testTitle",
		Output: "stdout",
		Level:  tc.LogLevel,
	}

	testLogger := NewLogger(&c)
	ctx := WithLogger(context.Background(), testLogger.WithField("aaa", "bbb"))
	ctxLogger := GetLogger(ctx)

	ctxLogger.Info(tc.LogLevel, "Info text")
	ctxLogger.Warning(tc.LogLevel, " Warn text")
	ctxLogger.Error(tc.LogLevel, "Err text")
	ctxLogger.Debug(tc.LogLevel, "Debug text")

}
func TestCreateStdoutJson(t *testing.T) {
	for _, tc := range generateLogs(t) {
		c := Config{
			Title:     "testTitle",
			Formatter: "json",
			Output:    "stdout",
			Level:     tc.LogLevel,
		}
		testLogger := NewLogger(&c)
		testLogger.Info(tc.LogLevel, "Info text")
		testLogger.Warning(tc.LogLevel, " Warn text")
		testLogger.Error(tc.LogLevel, "Err text")
		testLogger.Debug(tc.LogLevel, "Debug text")
	}
}

func TestSentry(t *testing.T) {
	for _, tc := range generateLogs(t) {
		c := Config{
			Output: "stdout",
			Title:  "testTitle",
			Level:  tc.LogLevel,
			Hooks: Hooks{
				Sentry: &Sentry{
					Tags: map[string]string{
						"site": "dev",
					},
					Level:         ErrorLevel,
					Async:         false,
					Timeout:       time.Second * 2,
					SSLSkipVerify: true,
					// TODO insert DNS for raven
					DSN: "",
				},
			},
		}
		testLogger := NewLogger(&c)
		testLogger.Info(tc.LogLevel, "Info text")
		testLogger.Warning(tc.LogLevel, " Warn text")
		testLogger.Error(tc.LogLevel, "Err text")
		testLogger.Debug(tc.LogLevel, "Debug text")
	}

}

func TestMarshal(t *testing.T) {
	c := Config{
		Output: "stdout",
		Title:  "testTitle",
		Level:  InfoLevel,
	}
	yamlData, err := yaml.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%s\n", yamlData)
	jsonlData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%s\n", jsonlData)
}

func TestUnMarshalYaml(t *testing.T) {
	data := `
title: testTitle
output: stdout
formatter: json
level: info
`
	var c Config
	err := yaml.Unmarshal([]byte(data), &c)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", c)
}
func TestUnMarshalJson(t *testing.T) {
	data := `
{
  "title": "testTitle",
  "output": "stdout",
  "formatter": "json",
  "level": "info"
}
`
	var c Config
	err := json.Unmarshal([]byte(data), &c)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", c)
}
