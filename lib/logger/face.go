package logger

import (
	"context"
	"io"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	WithField(key string, value interface{}) *logrus.Entry
	WithFields(fields logrus.Fields) *logrus.Entry
	WithError(err error) *logrus.Entry

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Tracef(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
	Trace(args ...interface{})

	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Println(args ...interface{})
	Warnln(args ...interface{})
	Warningln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})
	Traceln(args ...interface{})

	Log(level logrus.Level, args ...interface{})
	Writer() *io.PipeWriter
}

type ctxLog struct{}

func WithLogger(ctx context.Context, lg Logger) context.Context {
	return context.WithValue(ctx, ctxLog{}, lg)
}

func Gist(ctx context.Context) Logger {
	l, ok := ctx.Value(ctxLog{}).(Logger)
	if !ok {
		return logInstance
	}
	return l
}
