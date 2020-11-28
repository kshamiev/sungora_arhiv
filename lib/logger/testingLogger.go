package logger

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"reflect"
	"testing"
)

type testingWrapper struct {
	t      testing.TB
	mapper func(t testing.TB, level Level, args ...interface{})
	data   Fields
	err    string
}

var _ Logger = testingWrapper{}

func NewTesting(t testing.TB, mapper func(t testing.TB, level Level, args ...interface{}), data Fields, err string) Logger {
	return testingWrapper{t: t, mapper: mapper, data: data, err: err}
}

func defaultMapper(t testing.TB, level Level, args ...interface{}) {
	args = append([]interface{}{level}, args...)
	t.Log(args...)
}

func (tw testingWrapper) Log(level Level, args ...interface{}) {
	if tw.mapper == nil {
		defaultMapper(tw.t, level, args...)
	} else {
		tw.mapper(tw.t, level, args...)
	}
}

func (tw testingWrapper) CreateStdLogger() *log.Logger {
	return log.New(
		tw.Writer(),
		"",
		0)
}

// Level calls
func (tw testingWrapper) Trace(args ...interface{}) {
	tw.Log(TraceLevel, args...)
}

func (tw testingWrapper) Debug(args ...interface{}) {
	tw.Log(DebugLevel, args...)
}

func (tw testingWrapper) Info(args ...interface{}) {
	tw.Log(InfoLevel, args...)
}

func (tw testingWrapper) Warning(args ...interface{}) {
	tw.Log(WarnLevel, args...)
}

func (tw testingWrapper) Error(args ...interface{}) {
	tw.Log(TraceLevel, args...)
}

func (tw testingWrapper) Fatal(args ...interface{}) {
	tw.Log(FatalLevel, args...)
}

func (tw testingWrapper) Panic(args ...interface{}) {
	tw.Log(PanicLevel, args...)
}

func (tw testingWrapper) Panicln(args ...interface{}) {
	tw.Log(PanicLevel, args...)
}

func (tw testingWrapper) Audit(objectIdentifier, userIdentifier, action string) {
	tw.WithField("audit", true).
		WithField("log_type", "audit").
		WithField("object_id", objectIdentifier).
		WithField("user_id", userIdentifier).
		Log(WarnLevel, action)
}

func (tw testingWrapper) Writer() *io.PipeWriter {
	return tw.WriterLevel(InfoLevel)
}

func (tw testingWrapper) WriterLevel(level Level) *io.PipeWriter {
	reader, writer := io.Pipe()

	go tw.writerScanner(reader, level)

	return writer
}

func (tw testingWrapper) writerScanner(reader *io.PipeReader, level Level) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		tw.mapper(tw.t, level, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		tw.Error("Error while reading from Writer: ", err)
	}
	reader.Close()
}

func (tw testingWrapper) WithField(key string, value interface{}) Logger {
	return tw.WithFields(Fields{key: value})
}

func (tw testingWrapper) WithFields(fields Fields) Logger {
	data := make(Fields, len(tw.data)+len(fields))
	for k, v := range tw.data {
		data[k] = v
	}
	fieldErr := tw.err
	for k, v := range fields {
		isErrField := false
		if t := reflect.TypeOf(v); t != nil {
			switch t.Kind() {
			case reflect.Func, reflect.Chan:
				isErrField = true
			case reflect.Ptr:
				isErrField = t.Elem().Kind() == reflect.Func || t.Elem().Kind() == reflect.Chan
			}
		}
		if isErrField {
			tmp := fmt.Sprintf("can not add field %q", k)
			if fieldErr != "" {
				fieldErr = tw.err + ", " + tmp
			} else {
				fieldErr = tmp
			}
		} else {
			data[k] = v
		}
	}
	return &testingWrapper{t: tw.t, mapper: tw.mapper, data: data, err: fieldErr}
}

func (tw testingWrapper) WithError(err error) Logger {
	return tw.WithFields(Fields{ErrorField: err})
}
