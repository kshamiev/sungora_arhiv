package logger

import (
	"context"
	"io"
	"os"

	"sungora/lib/logger/graylog"
	"sungora/lib/typ"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	logrus.Ext1FieldLogger
	Log(level logrus.Level, args ...interface{})
	Logf(level logrus.Level, format string, args ...interface{})
	Writer() *io.PipeWriter
}

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
		inst.SetOutput(io.Discard)
	case outStdout:
		inst.SetOutput(os.Stdout)
	case outStderr:
		inst.SetOutput(os.Stderr)
	default:
		fp, err := os.OpenFile(config.Output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
		if err != nil {
			inst.SetOutput(os.Stdout)
			inst.Fatal(err)
		} else {
			inst.SetOutput(fp)
		}
	}

	if config.Hooks.Graylog.DSN != "" {
		hook := graylog.NewGraylogHook(config.Hooks.Graylog.DSN, nil)
		if config.Hooks.Graylog.Host != "" {
			hook.Host = config.Hooks.Graylog.Host
		}
		hook.Blacklist(config.Hooks.Graylog.Blacklist)
		inst.AddHook(hook)
	}

	instance = inst
	return instance
}

func Set(lg Logger) {
	instance = lg
}

type ctxLog struct{}

func WithLogger(ctx context.Context, lg Logger) context.Context {
	return context.WithValue(ctx, ctxLog{}, lg)
}

func Get(ctx context.Context) Logger {
	l, ok := ctx.Value(ctxLog{}).(Logger)
	if ok {
		return l
	}
	return instance
}
