package errs

import (
	"database/sql"
	"errors"
	"net/http"
)

// NewUnauthorized new error type
func NewUnauthorized(err error, args ...interface{}) *Errs {
	if err == nil {
		err = errors.New(http.StatusText(http.StatusUnauthorized))
	}
	return &Errs{
		codeHTTP: http.StatusUnauthorized,
		err:      err,
		kind:     trace(2),
		message:  getMessage(args),
	}
}

// NewNotFound new error type
func NewNotFound(err error, args ...interface{}) *Errs {
	if err == nil {
		err = errors.New(http.StatusText(http.StatusNotFound))
	}
	return &Errs{
		codeHTTP: http.StatusNotFound,
		err:      err,
		kind:     trace(2),
		message:  getMessage(args),
	}
}

// NewForbidden new error type
func NewForbidden(err error, args ...interface{}) *Errs {
	if err == nil {
		err = errors.New(http.StatusText(http.StatusForbidden))
	}

	return &Errs{
		codeHTTP: http.StatusForbidden,
		err:      err,
		kind:     trace(2),
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

	return &Errs{
		codeHTTP: codeHTTP,
		err:      err,
		kind:     trace(2),
		message:  getMessage(args),
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
