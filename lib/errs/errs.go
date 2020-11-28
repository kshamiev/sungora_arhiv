package errs

import (
	"errors"
	"net/http"

	"sungora/lib/app"
)

var (
	ErrNotFound            = errors.New(http.StatusText(http.StatusNotFound))
	ErrUnauthorized        = errors.New(http.StatusText(http.StatusUnauthorized))
	ErrForbidden           = errors.New(http.StatusText(http.StatusForbidden))
	ErrBadRequest          = errors.New(http.StatusText(http.StatusBadRequest))
	ErrInternalServerError = errors.New(http.StatusText(http.StatusInternalServerError))
)

// NewUnauthorized new error type
func NewUnauthorized(err error, msg ...string) *Errs {
	if err == nil {
		err = ErrUnauthorized
	}
	return &Errs{
		codeHTTP: http.StatusUnauthorized,
		err:      err,
		kind:     app.Trace(2),
		message:  getMessage(msg),
	}
}

// NewNotFound new error type
func NewNotFound(err error, msg ...string) *Errs {
	if err == nil {
		err = ErrNotFound
	}
	return &Errs{
		codeHTTP: http.StatusNotFound,
		err:      err,
		kind:     app.Trace(2),
		message:  getMessage(msg),
	}
}

// NewForbidden new error type
func NewForbidden(err error, msg ...string) *Errs {
	if err == nil {
		err = ErrForbidden
	}
	return &Errs{
		codeHTTP: http.StatusForbidden,
		err:      err,
		kind:     app.Trace(2),
		message:  getMessage(msg),
	}
}

// NewBadRequest new error type
func NewBadRequest(err error, msg ...string) *Errs {
	if err == nil {
		err = ErrBadRequest
	}
	return &Errs{
		codeHTTP: http.StatusBadRequest,
		err:      err,
		kind:     app.Trace(2),
		message:  getMessage(msg),
	}
}

// NewInternalServerError new error type
func NewInternalServerError(err error, msg ...string) *Errs {
	if err == nil {
		err = ErrInternalServerError
	}
	return &Errs{
		codeHTTP: http.StatusInternalServerError,
		err:      err,
		kind:     app.Trace(2),
		message:  getMessage(msg),
	}
}

type Errs struct {
	codeHTTP int    // код http
	err      error  // сама ошибка от внешнего сервиса или либы
	kind     string // где произошла ошибка
	message  string // сообщение для пользователя
}

// HTTPCode http status response
func (e *Errs) HTTPCode() int {
	return e.codeHTTP
}

// Error response advanced message to logs
func (e *Errs) Error() string {
	var k string
	if e.kind != "" {
		k = "; " + e.kind
	}
	if e.err != nil {
		return e.err.Error() + k
	}
	return http.StatusText(e.codeHTTP) + k
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
