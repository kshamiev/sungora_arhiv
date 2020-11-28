package errs

import (
	"errors"
	"strings"
	"testing"

	"google.golang.org/grpc/codes"
)

func TestErrsGRPC(t *testing.T) {
	err := NewGRPC(codes.NotFound, errors.New("LIBRARY ERROR"), "user error")
	e := ParseGRPC(err)

	const LibraryError = "LIBRARY ERROR"

	switch {
	case !strings.Contains(e.Error(), LibraryError):
		t.Error("parameter: Error")
	case e.Response() != "user error":
		t.Error("parameter: Response")
	case e.codeHTTP != 404:
		t.Error("parameter: codeHTTP")
	}
}
