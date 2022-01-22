package errs

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"sungora/lib/app"
)

// NewUnauthorized new error type
func NewUnauthorized(err error, args ...interface{}) *Errs {
	if err == nil {
		err = errors.New(http.StatusText(http.StatusUnauthorized))
	}
	t, _, _, _ := app.Trace(2)
	return &Errs{
		codeHTTP: http.StatusUnauthorized,
		err:      err,
		kind:     t,
		message:  getMessage(args),
	}
}

// NewNotFound new error type
func NewNotFound(err error, args ...interface{}) *Errs {
	if err == nil {
		err = errors.New(http.StatusText(http.StatusNotFound))
	}
	t, _, _, _ := app.Trace(2)
	return &Errs{
		codeHTTP: http.StatusNotFound,
		err:      err,
		kind:     t,
		message:  getMessage(args),
	}
}

// NewForbidden new error type
func NewForbidden(err error, args ...interface{}) *Errs {
	if err == nil {
		err = errors.New(http.StatusText(http.StatusForbidden))
	}
	t, _, _, _ := app.Trace(2)
	return &Errs{
		codeHTTP: http.StatusForbidden,
		err:      err,
		kind:     t,
		message:  getMessage(args),
	}
}

// NewBadRequest new error type
func NewBadRequest(err error, args ...interface{}) *Errs {
	if err == nil {
		err = errors.New(http.StatusText(http.StatusBadRequest))
	}
	codeHTTP := http.StatusBadRequest
	if sql.ErrNoRows == err {
		codeHTTP = http.StatusNotFound
	}
	t, _, _, _ := app.Trace(2)
	return &Errs{
		codeHTTP: codeHTTP,
		err:      err,
		kind:     t,
		message:  getMessage(args),
		trace:    app.Traces(),
	}
}

// New custom error application
func New(err error, msg Message, args ...interface{}) *Errs {
	if err == nil {
		err = msg.Error(args...)
	}
	codeHTTP := http.StatusBadRequest
	if sql.ErrNoRows == err {
		codeHTTP = http.StatusNotFound
	}
	t, _, _, _ := app.Trace(2)
	return &Errs{
		codeHTTP: codeHTTP,
		err:      err,
		kind:     t,
		message:  msg.String(args...),
		trace:    app.Traces(),
	}
}

// ////

type Errs struct {
	codeHTTP int    // код http
	err      error  // сама ошибка от внешнего сервиса или либы
	kind     string // где произошла ошибка
	message  string // сообщение для пользователя
	trace    []string
}

// HTTPCode http status response
func (e *Errs) HTTPCode() int {
	return e.codeHTTP
}

// Error response advanced message to logs
func (e *Errs) Error() string {
	if e.err != nil {
		return e.kind + " - " + e.err.Error()
	}
	return e.kind + " - " + http.StatusText(e.codeHTTP)
}

// Response response message to user
func (e *Errs) Response() string {
	if e.message != "" {
		return e.message
	} else if e.err != nil {
		return e.err.Error()
	}
	return http.StatusText(e.codeHTTP)
}

func (e *Errs) Trace() []string {
	return e.trace
}

func getMessage(args []interface{}) string {
	switch len(args) {
	case 0:
		return ""
	case 1:
		return args[0].(string)
	default:
		return fmt.Sprintf(args[0].(string), args[1:]...)
	}
}
