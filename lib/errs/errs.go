package errs

import (
	"database/sql"
	"net/http"
)

// NewUnauthorized new error type
func NewUnauthorized(err error, args ...interface{}) *Errs {
	e := New(err, args...)
	e.codeHTTP = http.StatusUnauthorized
	return e
}

// NewNotFound new error type
func NewNotFound(err error, args ...interface{}) *Errs {
	e := New(err, args...)
	e.codeHTTP = http.StatusUnauthorized
	return e
}

// NewForbidden new error type
func NewForbidden(err error, args ...interface{}) *Errs {
	e := New(err, args...)
	e.codeHTTP = http.StatusUnauthorized
	return e
}

// New custom error application
func New(err error, args ...interface{}) *Errs {
	var msg Message
	if len(args) > 0 {
		switch v := args[0].(type) {
		case Message:
			msg = v
			args = args[1:]
		case string:
			msg = Message(v)
			args = args[1:]
		}
	}
	if err == nil {
		err = msg.Error(args...)
	}
	codeHTTP := http.StatusBadRequest
	if sql.ErrNoRows == err {
		codeHTTP = http.StatusNotFound
	}
	t, _, _, _ := Trace(2)
	return &Errs{
		codeHTTP: codeHTTP,
		err:      err,
		kind:     t,
		message:  msg.String(args...),
		trace:    Traces(),
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
		return e.kind + " " + e.err.Error()
	}
	return e.kind + " " + http.StatusText(e.codeHTTP)
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
