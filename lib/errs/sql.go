package errs

import (
	"database/sql"
	"net/http"

	"sungora/lib/app"
)

func NewSQL(err error, msg ...string) error {
	e := &Errs{
		codeHTTP: http.StatusBadRequest,
		err:      err,
		kind:     app.Trace(2),
		message:  getMessage(msg),
	}
	if sql.ErrNoRows == err {
		e.codeHTTP = http.StatusNotFound
	}
	return e
}

func getMessage(s []string) string {
	if len(s) > 0 {
		return s[0]
	}
	return ""
}
