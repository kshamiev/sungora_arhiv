package request

import (
	"context"
	"errors"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"sungora/lib/enum"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/cors"
	swaggerFiles "github.com/swaggo/files"
	"google.golang.org/grpc/metadata"

	"sungora/lib/errs"
	"sungora/lib/response"
	"sungora/lib/typ"
)

type Mid struct {
	token      string
	signingKey string
	dirStatic  string
}

func NewMid(token, signingKey, dirStatic string) *Mid {
	return &Mid{
		token:      token,
		signingKey: signingKey,
		dirStatic:  dirStatic,
	}
}

// Authentication аутентификация по токену из куки
func (mid *Mid) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := response.New(r, w)

		// получаем токен
		token := rw.CookieGet(mid.token)
		if token == "" {
			rw.JSONError(errs.NewUnauthorized(nil, "token is empty"))
			return
		}

		// анализируем токен
		us, err := mid.VerifyToken(token)
		if err != nil {
			rw.JSONError(errs.NewForbidden(err, "token is invalid"))
			return
		}

		ctx := r.Context()
		ctx = metadata.AppendToOutgoingContext(ctx, string(response.CtxToken), token)
		ctx = context.WithValue(ctx, response.CtxToken, token)
		ctx = context.WithValue(ctx, response.CtxUser, us)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Cors добавление заголовка ConfigCors
func (mid *Mid) Cors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPatch,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
}

// Static статика или отдача существующего файла по запросу
func (mid *Mid) Static(p string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.New(r, w)
		rw.Static(p + r.URL.Path)
	}
}

// GenToken генерация jwt токена с данными по соли и установка его таймаута
func (mid *Mid) GenToken(us *response.User, dur time.Duration) (token string, err error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": us.ID.String(),
		"login":  us.Login,
		"roles":  us.Roles,
		"exp":    time.Now().Add(dur).Unix(),
	}).SignedString([]byte(mid.signingKey))
}

// VerifyToken проверка jwt токена по соли и его таймаута, получение данных
func (mid *Mid) VerifyToken(token string) (*response.User, error) {
	var tokenObj *jwt.Token
	var err error

	if tokenObj, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(mid.signingKey), nil
	}); err != nil {
		return nil, err
	}

	if claims, ok := tokenObj.Claims.(jwt.MapClaims); ok && tokenObj.Valid {
		if claims["exp"] == nil {
			return nil, errors.New("bad token, re login please")
		}

		_, ok := claims["exp"].(float64)
		if !ok {
			return nil, errors.New("error get tiken exp")
		}

		uid := typ.UUIDNew()
		if err := uid.Scan(claims["userID"].(string)); err != nil {
			return nil, err
		}
		us := &response.User{ID: uid}

		if _, ok := claims["login"].(string); !ok {
			return nil, errors.New("error get login")
		}
		us.Login = claims["login"].(string)

		if _, ok := claims["roles"].([]interface{}); !ok {
			return nil, errors.New("error get roles")
		}
		us.Roles = make([]enum.Role, len(claims["roles"].([]interface{})))
		for i, role := range claims["roles"].([]interface{}) {
			us.Roles[i] = enum.Role(role.(string))
		}
		return us, nil
	}
	return nil, errors.New("error get tokenObj.Claims")
}

// документация и и тестирование api приложения
func (mid *Mid) Swagger(docDefault string) http.HandlerFunc {
	type Config struct {
		URL          string
		DeepLinking  bool
		DocExpansion string
		DomID        string
	}
	cfg := &Config{
		URL:          docDefault,
		DeepLinking:  true,
		DocExpansion: "list",
		DomID:        "#swagger-ui",
	}

	t := template.New("swagger_index.html")
	index, _ := t.Parse(indexTpl)

	var re = regexp.MustCompile(`^(.*/)([^?].*)?[?|.]*$`)

	return func(w http.ResponseWriter, r *http.Request) {
		matches := re.FindStringSubmatch(r.RequestURI)
		pathUri := matches[2]
		prefix := matches[1]

		h := swaggerFiles.Handler
		h.Prefix = prefix

		if strings.Contains(pathUri, ".json") {
			data, err := ioutil.ReadFile(mid.dirStatic + "/swagger/" + pathUri)
			if err != nil {
				_, _ = w.Write([]byte(err.Error()))
				return
			}
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			_, _ = w.Write(data)
			return
		}

		switch pathUri {
		case "index.html":
			_ = index.Execute(w, cfg)
		case "":
			http.Redirect(w, r, prefix+"index.html", 301)
		default:
			h.ServeHTTP(w, r)
		}
	}
}

// nolint
const indexTpl = `<!-- HTML for static distribution bundle build -->
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Swagger UI</title>
  <link href="https://fonts.googleapis.com/css?family=Open+Sans:400,700|Source+Code+Pro:300,600|Titillium+Web:400,600,700" rel="stylesheet">
  <link rel="stylesheet" type="text/css" href="./swagger-ui.css" >
  <link rel="icon" type="image/png" href="./favicon-32x32.png" sizes="32x32" />
  <link rel="icon" type="image/png" href="./favicon-16x16.png" sizes="16x16" />
  <style>
    html
    {
        box-sizing: border-box;
        overflow: -moz-scrollbars-vertical;
        overflow-y: scroll;
    }
    *,
    *:before,
    *:after
    {
        box-sizing: inherit;
    }

    body {
      margin:0;
      background: #fafafa;
    }
  </style>
</head>

<body>

<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" style="position:absolute;width:0;height:0">
  <defs>
    <symbol viewBox="0 0 20 20" id="unlocked">
      <path d="M15.8 8H14V5.6C14 2.703 12.665 1 10 1 7.334 1 6 2.703 6 5.6V6h2v-.801C8 3.754 8.797 3 10 3c1.203 0 2 .754 2 2.199V8H4c-.553 0-1 .646-1 1.199V17c0 .549.428 1.139.951 1.307l1.197.387C5.672 18.861 6.55 19 7.1 19h5.8c.549 0 1.428-.139 1.951-.307l1.196-.387c.524-.167.953-.757.953-1.306V9.199C17 8.646 16.352 8 15.8 8z"></path>
    </symbol>

    <symbol viewBox="0 0 20 20" id="locked">
      <path d="M15.8 8H14V5.6C14 2.703 12.665 1 10 1 7.334 1 6 2.703 6 5.6V8H4c-.553 0-1 .646-1 1.199V17c0 .549.428 1.139.951 1.307l1.197.387C5.672 18.861 6.55 19 7.1 19h5.8c.549 0 1.428-.139 1.951-.307l1.196-.387c.524-.167.953-.757.953-1.306V9.199C17 8.646 16.352 8 15.8 8zM12 8H8V5.199C8 3.754 8.797 3 10 3c1.203 0 2 .754 2 2.199V8z"/>
    </symbol>

    <symbol viewBox="0 0 20 20" id="close">
      <path d="M14.348 14.849c-.469.469-1.229.469-1.697 0L10 11.819l-2.651 3.029c-.469.469-1.229.469-1.697 0-.469-.469-.469-1.229 0-1.697l2.758-3.15-2.759-3.152c-.469-.469-.469-1.228 0-1.697.469-.469 1.228-.469 1.697 0L10 8.183l2.651-3.031c.469-.469 1.228-.469 1.697 0 .469.469.469 1.229 0 1.697l-2.758 3.152 2.758 3.15c.469.469.469 1.229 0 1.698z"/>
    </symbol>

    <symbol viewBox="0 0 20 20" id="large-arrow">
      <path d="M13.25 10L6.109 2.58c-.268-.27-.268-.707 0-.979.268-.27.701-.27.969 0l7.83 7.908c.268.271.268.709 0 .979l-7.83 7.908c-.268.271-.701.27-.969 0-.268-.269-.268-.707 0-.979L13.25 10z"/>
    </symbol>

    <symbol viewBox="0 0 20 20" id="large-arrow-down">
      <path d="M17.418 6.109c.272-.268.709-.268.979 0s.271.701 0 .969l-7.908 7.83c-.27.268-.707.268-.979 0l-7.908-7.83c-.27-.268-.27-.701 0-.969.271-.268.709-.268.979 0L10 13.25l7.418-7.141z"/>
    </symbol>

    <symbol viewBox="0 0 24 24" id="jump-to">
      <path d="M19 7v4H5.83l3.58-3.59L8 6l-6 6 6 6 1.41-1.41L5.83 13H21V7z"/>
    </symbol>

    <symbol viewBox="0 0 24 24" id="expand">
      <path d="M10 18h4v-2h-4v2zM3 6v2h18V6H3zm3 7h12v-2H6v2z"/>
    </symbol>
  </defs>
</svg>

<div id="swagger-ui"></div>

<script src="./swagger-ui-bundle.js"> </script>
<script src="./swagger-ui-standalone-preset.js"> </script>
<script>
window.onload = function() {
  // Build a system
  const ui = SwaggerUIBundle({
    url: "{{.URL}}",
    deepLinking: {{.DeepLinking}},
    docExpansion: "{{.DocExpansion}}",
    dom_id: "{{.DomID}}",
    validatorUrl: null,
    presets: [
      SwaggerUIBundle.presets.apis,
      SwaggerUIStandalonePreset
    ],
    plugins: [
      SwaggerUIBundle.plugins.DownloadUrl
    ],
    layout: "StandaloneLayout"
  })

  window.ui = ui
}
</script>
</body>

</html>
`
