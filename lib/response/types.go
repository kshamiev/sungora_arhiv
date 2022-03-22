package response

import (
	"sungora/lib/enum"
	"sungora/lib/typ"
)

type CtxKey string

const (
	IndexHtml = "index.html"

	CtxUser  CtxKey = "user"
	CtxToken CtxKey = "token"
)

type User struct {
	ID    typ.UUID
	Login string
	Roles []enum.Role
}

type Error interface {
	Error() string
	Response() string
	Trace() []string
	HTTPCode() int
}

type Data struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
