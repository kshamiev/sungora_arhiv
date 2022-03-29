package response

type CtxKey string

const (
	IndexHtml = "index.html"

	CtxUser  CtxKey = "user"
	CtxToken CtxKey = "token"
)

type User struct {
	ID    int64
	Login string
	Roles []string
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
