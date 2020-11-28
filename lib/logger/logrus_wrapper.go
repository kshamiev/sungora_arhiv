package logger

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"sync/atomic"

	"github.com/sirupsen/logrus"
)

// logrusLevelMap
var logrusLevelMap = map[Level]logrus.Level{
	emptyLevel: logrus.TraceLevel,
	PanicLevel: logrus.PanicLevel,
	FatalLevel: logrus.FatalLevel,
	ErrorLevel: logrus.ErrorLevel,
	WarnLevel:  logrus.WarnLevel,
	InfoLevel:  logrus.InfoLevel,
	DebugLevel: logrus.DebugLevel,
	TraceLevel: logrus.TraceLevel,
}

func newLogrusWrapper(config *Config) Logger {
	logger := logrus.New()
	audit := logrus.New()
	if config == nil || (*config == Config{}) {
		logger.SetOutput(os.Stdout)
		logger.SetLevel(logrusLevelMap[TraceLevel])
		logger.SetFormatter(&logrus.TextFormatter{})
		audit.SetOutput(os.Stdout)
		audit.SetLevel(logrus.WarnLevel)
		audit.SetFormatter(&logrus.JSONFormatter{})
		return &logrusWrapper{
			done:  1,
			Entry: logger.WithFields(nil),
			audit: audit.WithFields(nil),
		}
	}
	logger.SetReportCaller(config.ReportCaller)

	var fl, fa io.WriteCloser

	switch config.Output {
	case Stdout, "":
		logger.SetOutput(os.Stdout)
	case Stderr:
		logger.SetOutput(os.Stderr)
	case Vacuum:
		logger.SetOutput(ioutil.Discard)
	default:
		var err error
		fl, err = os.OpenFile(config.Output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			logger.SetOutput(os.Stdout)
			logger.WithError(err).Debug("cant create log file, falling to stdout")
		} else {
			logger.SetOutput(fl)
		}
	}

	formatter, ok := formatters[config.Formatter]
	if !ok {
		formatter = &logrus.TextFormatter{}
	}
	logger.SetFormatter(formatter)
	logger.SetLevel(logrusLevelMap[config.Level])

	addHooks(logger, config)

	switch config.AuditOutput {
	case Stderr:
		audit.SetOutput(os.Stderr)
	case Stdout, "":
		audit.SetOutput(os.Stdout)
	case Vacuum:
		audit.SetOutput(ioutil.Discard)
	default:
		var err error
		fa, err = os.OpenFile(config.Output, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			audit.SetOutput(os.Stdout)
			audit.WithError(err).Debug("cant create audit log file, falling to stdout")
		} else {
			audit.SetOutput(fa)
		}
	}

	if config.Title != "" {
		return &logrusWrapper{
			done:  1,
			Entry: logger.WithField(TitleField, config.Title),
			audit: audit.WithField(TitleField, config.Title),
			fl:    fl,
			fa:    fa,
		}
	}
	return &logrusWrapper{
		done:  1,
		Entry: logger.WithFields(nil),
		audit: audit.WithFields(nil),
		fl:    fl,
		fa:    fa,
	}
}

func addHooks(logger *logrus.Logger, config *Config) {
	if config.Hooks.Sentry != nil {
		sh, err := sentryHook(config.Hooks.Sentry)
		if err == nil {
			logger.AddHook(sh)
		} else {
			logger.WithError(err).Debug("can't add hook sentry")
		}
	}
	//nolint:gocritic
	//if config.Hooks.Syslog != nil {
	//	sh, err := sysloggerHook(config.Hooks.Syslog)
	//	if err == nil {
	//		logger.AddHook(sh)
	//	} else {
	//		logger.WithError(err).Debug("can't add hook syslog")
	//	}
	// }
	//if config.Hooks.FileName != nil { //  неизвестный баг, но в месте с этим хуком виснет весь логгер (ну и аппка)
	//	logger.AddHook(filenameHook(config.Hooks.FileName))
	//}
}

// logrusWrapper is a local wrapper around logrus
type logrusWrapper struct {
	*logrus.Entry
	audit  *logrus.Entry
	done   uint32
	m      sync.Mutex
	fl, fa io.WriteCloser
}

func (w *logrusWrapper) Close() {
	if w.fl != nil {
		w.fl.Close()
	}
	if w.fa != nil {
		w.fa.Close()
	}
}

func (w *logrusWrapper) Do(f func()) {
	if atomic.LoadUint32(&w.done) == 0 {
		w.doSlow(f)
	}
}

func (w *logrusWrapper) doSlow(f func()) {
	w.m.Lock()
	defer w.m.Unlock()
	if w.done == 0 {
		defer atomic.StoreUint32(&w.done, 1)
		f()
	}
}

func (w *logrusWrapper) initDefault() {
	if w.Entry == nil {
		w.Entry = &logrus.Entry{
			Logger: &logrus.Logger{
				Out:          os.Stdout,
				Formatter:    new(logrus.TextFormatter),
				Hooks:        make(logrus.LevelHooks),
				Level:        logrus.InfoLevel,
				ExitFunc:     os.Exit,
				ReportCaller: false,
			},
			Data: make(logrus.Fields, 6),
		}
	}
	if w.audit == nil {
		w.audit = &logrus.Entry{
			Logger: &logrus.Logger{
				Out:          os.Stdout,
				Formatter:    new(logrus.JSONFormatter),
				Hooks:        make(logrus.LevelHooks),
				Level:        logrus.WarnLevel,
				ExitFunc:     os.Exit,
				ReportCaller: false,
			},
			Data: make(logrus.Fields, 6),
		}
	}
}

func (w *logrusWrapper) Writer() *io.PipeWriter {
	w.Do(w.initDefault)
	return w.Entry.Writer()
}

// Log is a wrapper around logrus.Log
func (w *logrusWrapper) Log(level Level, args ...interface{}) {
	w.Do(w.initDefault)
	w.Entry.Log(logrusLevelMap[level], args...)
}

// Log is a wrapper around logrus.Log
func (w *logrusWrapper) Audit(objectIdentifier, userIdentifier, action string) {
	w.Do(w.initDefault)
	w.audit.
		WithField("log_type", "audit").
		WithField("object_id", objectIdentifier).
		WithField("user_id", userIdentifier).
		Log(logrus.WarnLevel, action)
}

// Log is a wrapper around logrus.Log
func (w *logrusWrapper) CreateStdLogger() *log.Logger {
	return log.New(
		w.Writer(),
		"",
		0)
}

// WithFields is a wrapper around logrus.WithFields
func (w *logrusWrapper) WithField(key string, value interface{}) Logger {
	w.Do(w.initDefault)
	return &logrusWrapper{done: w.done, Entry: w.Entry.WithField(key, value)}
}
func (w *logrusWrapper) WithFields(fields Fields) Logger {
	w.Do(w.initDefault)
	logrusFields := logrus.Fields(fields)
	return &logrusWrapper{done: w.done, Entry: w.Entry.WithFields(logrusFields)}
}

// WithError appends to logger 'error' key
func (w *logrusWrapper) WithError(err error) Logger {
	w.Do(w.initDefault)
	return w.WithFields(Fields{ErrorField: err})
}

// WrapLogrusEntry wrap logrus logger to our interface
func WrapLogrusEntry(l *logrus.Entry) (Logger, error) {
	if l != nil {
		return &logrusWrapper{done: 1, Entry: l}, nil
	}
	return nil, errors.New("can't wrap nil")
}

// WrapLogrusLogger wrap logrus entry to our interface
func WrapLogrusLogger(l *logrus.Logger) (Logger, error) {
	if l != nil {
		return &logrusWrapper{done: 1, Entry: l.WithField("wrapped", "wrapped")}, nil
	}
	return nil, errors.New("can't wrap nil")
}
