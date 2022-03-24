package errs

import "fmt"

type Message string

func (m Message) Error(args ...interface{}) error {
	if m == "" {
		return fmt.Errorf(fmt.Sprint(args...))
	}
	return fmt.Errorf(string(m), args...)
}
func (m Message) String(args ...interface{}) string {
	if m == "" {
		return fmt.Sprint(args...)
	}
	return fmt.Sprintf(string(m), args...)
}

const (
	UserOne Message = "Первая ошибка пользователя %d"
	UserTwo Message = "Вторая ошибка пользователя %s"
	// etc...
)
