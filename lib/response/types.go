package response

import (
	"sungora/lib/uuid"
	"sungora/src/typ"
)

type CtxKey string

const (
	CtxUser  CtxKey = "user"
	CtxToken CtxKey = "token"
)

type User struct {
	ID    uuid.UUID
	Login string
	Roles []typ.Role
}

// interface for responses with an error
type Error interface {
	Error() string
	Response() string
	HTTPCode() int
}

type Data struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
