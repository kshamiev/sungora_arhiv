package log

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

var lg = logrus.New().WithField(titleField, "default")

func Init(config *Config) *logrus.Entry {
	l := logrus.New()

	l.SetLevel(config.Level)

	switch config.Formatter {
	case formatterJSON:
		l.SetFormatter(&logrus.TextFormatter{})
	default:
		l.SetFormatter(&logrus.JSONFormatter{})
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

func WithLogger(ctx context.Context, logger *logrus.Entry) context.Context {
	return context.WithValue(ctx, ctxLog{}, logger)
}

func Gist(ctx context.Context) *logrus.Entry {
	l, ok := ctx.Value(ctxLog{}).(*logrus.Entry)
	if !ok {
		return lg
	}
	return l
}
