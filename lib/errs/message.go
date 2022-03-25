package errs

import (
	"fmt"
)

type Message string

func (m Message) Msg(args ...interface{}) string {
	if m == "" {
		return fmt.Sprint(args...)
	}
	return fmt.Sprintf(string(m), args...)
}

func (m Message) String() string {
	return string(m)
}
