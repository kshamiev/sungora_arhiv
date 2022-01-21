package logger

import (
	"context"
	"os"

	"sungora/lib/typ"

	"github.com/sirupsen/logrus"
)

var lg = logrus.New().WithField(titleField, "default")

func Init(config *Config) Logger {
	l := logrus.New()
	l.SetLevel(config.Level)
	switch config.Formatter {
	case formatterJSON:
		l.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat:   typ.TimeFormatGMDHIS,
			DisableTimestamp:  false,
			DisableHTMLEscape: false,
			DataKey:           "",
			FieldMap:          nil,
			CallerPrettyfier:  nil,
			PrettyPrint:       false,
		})
	default:
		l.SetFormatter(&logrus.TextFormatter{})
	}
	switch config.Output {
	case stdout:
		l.SetOutput(os.Stdout)
	default:
		fp, err := os.OpenFile(config.Output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			l.SetOutput(os.Stdout)
			l.Fatal(err)
		} else {
			l.SetOutput(fp)
		}
	}
	lg = l.WithField(titleField, config.Title)
	return lg
}

// ////

type ctxLog struct{}

func WithLogger(ctx context.Context, lg Logger) context.Context {
	return context.WithValue(ctx, ctxLog{}, lg)
}

func Gist(ctx context.Context) Logger {
	l, ok := ctx.Value(ctxLog{}).(Logger)
	if !ok {
		return lg
	}
	return l
}
