package errs

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrs(t *testing.T) {
	err := FunctionLevel2()
	if e, ok := err.(*Errs); ok {
		fmt.Println("UI error: " + e.Response())
		fmt.Println(e.Error())
		for _, l := range e.Trace() {
			fmt.Println(l)
		}
	}

	err = FunctionLevel22()
	if e, ok := err.(*Errs); ok {
		fmt.Println("UI error: " + e.Response())
		fmt.Println(e.Error())
		for _, l := range e.Trace() {
			fmt.Println(l)
		}
	}
}

func FunctionLevel2() error {
	return FunctionLevel3()
}

func FunctionLevel3() error {
	return FunctionLevel4()
}

func FunctionLevel4() error {
	return NewBadRequest(errors.New("focus pocus"), "decimal: %d, string: %s, float", 34, "popcorn")
}

func FunctionLevel22() error {
	return FunctionLevel33()
}

func FunctionLevel33() error {
	return FunctionLevel44()
}

func FunctionLevel44() error {
	return New(nil, UserOne, 567)
}
