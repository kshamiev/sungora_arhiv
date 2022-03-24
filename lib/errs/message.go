package errs

import (
	"errors"
	"fmt"
)

type Message string

func (m Message) New(args ...interface{}) error {
	if m == "" {
		return errors.New(fmt.Sprint(args...))
	}
	return fmt.Errorf(string(m), args...)
}
func (m Message) Msg(args ...interface{}) string {
	if m == "" {
		return fmt.Sprint(args...)
	}
	return fmt.Sprintf(string(m), args...)
}

func (m Message) String() string {
	return string(m)
}
