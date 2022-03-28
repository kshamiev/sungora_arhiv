package errs

import (
	"errors"
	"testing"
)

const (
	ErrOne    Message = "Первая ошибка пользователя %d"
	ErrSample Message = "decimal: %d, string: %s, float %f"
)

func TestErrs(t *testing.T) {
	err := FunctionLevel2()
	if e, ok := err.(*Errs); ok {
		t.Log("(front) UI error: " + e.Response())
		t.Log(e.Error())
		for _, l := range e.Trace() {
			t.Log(l)
		}
	}

	err = FunctionLevel22()
	if e, ok := err.(*Errs); ok {
		t.Log("(front) UI error: " + e.Response())
		t.Log(e.Error())
		for _, l := range e.Trace() {
			t.Log(l)
		}
	}

	t.Log(ErrSample.Msg(34, "popcorn", 45.78))

	t.Log(ErrSample)
}

func FunctionLevel2() error {
	return FunctionLevel3()
}

func FunctionLevel3() error {
	return FunctionLevel4()
}

func FunctionLevel4() error {
	return New(errors.New("focus pocus"), ErrSample, 34, "popcorn", 45.78)
}

func FunctionLevel22() error {
	return FunctionLevel33()
}

func FunctionLevel33() error {
	return FunctionLevel44()
}

func FunctionLevel44() error {
	return New(nil, ErrOne, 567)
}
