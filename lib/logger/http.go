package logger

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"google.golang.org/grpc/metadata"
)

type ContextKey string

const LogTraceID = "trace-id"
const LogTraceAPI = "api"

const CtxTraceID ContextKey = "trace-id"
const CtxTraceAPI ContextKey = "api"

func Middleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := uuid.New().String()
			ctx := r.Context()

			lg := Gist(ctx).WithField(LogTraceID, requestID)
			ctx = WithLogger(ctx, lg)
			ctx = boil.WithDebugWriter(ctx, lg.Writer())
			ctx = context.WithValue(ctx, CtxTraceID, requestID)
			ctx = metadata.AppendToOutgoingContext(ctx, LogTraceID, requestID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
