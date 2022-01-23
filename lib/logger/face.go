package logger

import (
	"context"
	"io"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	logrus.Ext1FieldLogger
	Log(level logrus.Level, args ...interface{})
	Logf(level logrus.Level, format string, args ...interface{})
	Writer() *io.PipeWriter
}

type ctxLog struct{}

func WithLogger(ctx context.Context, lg Logger) context.Context {
	return context.WithValue(ctx, ctxLog{}, lg)
}

func Gist(ctx context.Context) Logger {
	l, ok := ctx.Value(ctxLog{}).(Logger)
	if ok {
		return l
	}
	return instance
}
