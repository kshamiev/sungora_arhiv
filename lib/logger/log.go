package logger

import (
	"io/ioutil"
	"os"

	"sungora/lib/logger/graylog"
	"sungora/lib/typ"

	"github.com/sirupsen/logrus"
)

var instance Logger = logrus.New()

func Init(config *Config) Logger {
	inst := logrus.New()
	inst.SetLevel(config.Level)
	inst.SetReportCaller(config.IsCaller)

	switch config.Formatter {
	case formatterJSON:
		inst.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:   typ.TimeFormatGMDHIS,
			DisableTimestamp:  false,
			DisableHTMLEscape: false,
			DataKey:           "",
			FieldMap:          nil,
			CallerPrettyfier:  nil,
			PrettyPrint:       false,
		})
	default:
		inst.SetFormatter(&logrus.TextFormatter{})
	}

	switch config.Output {
	case "", outEmpty:
		inst.SetOutput(ioutil.Discard)
	case outStdout:
		inst.SetOutput(os.Stdout)
	case outStderr:
		inst.SetOutput(os.Stderr)
	default:
		fp, err := os.OpenFile(config.Output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			inst.SetOutput(os.Stdout)
			inst.Fatal(err)
		} else {
			inst.SetOutput(fp)
		}
	}

	if config.Hooks.Graylog.DSN != "" {
		hook := graylog.NewGraylogHook(config.Hooks.Graylog.DSN, nil)
		hook.Host = "testing.local"
		hook.Blacklist([]string{})
		inst.AddHook(hook)
	}

	inst.WithField("param2", "fsdfdsgf").Warning("pupsik 1")

	instance = inst
	return instance
}

func SetCustomLogger(lg Logger) {
	instance = lg
}
