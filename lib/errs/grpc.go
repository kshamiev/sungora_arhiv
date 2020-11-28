package errs

import (
	"errors"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"sungora/lib/app"
)

const delimiter = "!===!"

// NewGRPC new error type
// Deprecated
func NewGRPC(code codes.Code, err error, msg string) error {
	if msg != "" {
		return status.Errorf(code, err.Error()+"; "+app.Trace(2)+delimiter+msg)
	}
	return status.Errorf(code, err.Error()+"; "+app.Trace(2))
}

// ParseGRPC parse error type
// Deprecated
func ParseGRPC(err error) *Errs {
	e := &Errs{}
	e.kind = app.Trace(2)

	if s, ok := status.FromError(err); ok {
		l := strings.Split(s.Message(), delimiter)
		e.codeHTTP = runtime.HTTPStatusFromCode(s.Code())
		e.err = errors.New(l[0])
		if len(l) > 1 {
			e.message = l[1]
		}
	}
	return e
}
