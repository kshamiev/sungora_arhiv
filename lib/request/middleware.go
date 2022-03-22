package request

import (
	"context"
	"errors"
	"net/http"
	"path"
	"strings"
	"time"

	"sungora/lib/enum"
	"sungora/lib/errs"
	"sungora/lib/logger"
	"sungora/lib/request/observability"
	"sungora/lib/response"
	"sungora/lib/typ"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Mid struct {
	token      string
	signingKey string
}

func NewMid(token, signingKey string) *Mid {
	return &Mid{
		token:      token,
		signingKey: signingKey,
	}
}

// Authentication аутентификация по токену из куки
func (mid *Mid) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := response.New(r, w)

		// получаем токен
		token := rw.CookieGet(mid.token)
		if token == "" {
			rw.JSON(errs.NewUnauthorized(nil, "token is empty"))
			return
		}

		// анализируем токен
		us, err := mid.VerifyToken(token)
		if err != nil {
			rw.JSON(errs.NewForbidden(err, "token is invalid"))
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
func (mid *Mid) Static(pathWww string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := response.New(r, w)
		rw.Static(pathWww + r.URL.Path)
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

func (mid *Mid) Observation() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return &ochttp.Handler{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				ctx := req.Context()
				// trick to calculate real route pattern for subrouters
				rctx := chi.RouteContext(ctx)
				nc := chi.Context{}
				rctx.Routes.Match(&nc, req.Method, req.RequestURI)

				httpPath := strings.ReplaceAll(path.Join(nc.RoutePatterns...), "/*/", "/")
				span := trace.FromContext(ctx)
				span.AddAttributes(trace.StringAttribute(ochttp.PathAttribute, httpPath))
				ochttp.SetRoute(ctx, httpPath)

				requestID := span.SpanContext().TraceID.String()
				lg := logger.Get(ctx).WithField(logger.TraceID, requestID)
				ctx = logger.WithLogger(ctx, lg)
				ctx = boil.WithDebugWriter(ctx, lg.Writer())
				ctx = context.WithValue(ctx, logger.CtxTraceID, requestID)
				ctx = metadata.AppendToOutgoingContext(ctx, logger.TraceID, requestID)

				next.ServeHTTP(w, req.WithContext(ctx))
			}),
			FormatSpanName: func(req *http.Request) string {
				rctx := chi.RouteContext(req.Context())
				nc := chi.Context{}
				if rctx.Routes != nil {
					rctx.Routes.Match(&nc, req.Method, req.RequestURI)
				}
				if rctx == nil {
					return ""
				}
				return strings.ReplaceAll(path.Join(nc.RoutePatterns...), "/*/", "/")
			},
			Propagation: observability.NewHTTPFormat(),
		}
	}
}

func Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
		resp interface{}, err error) {
		//
		md, ok := metadata.FromIncomingContext(ctx)
		if ok && md.Get(string(response.CtxToken)) != nil {
			ctx = context.WithValue(ctx, response.CtxToken, md.Get(string(response.CtxToken))[0])
		}
		return handler(ctx, req)
	}
}
