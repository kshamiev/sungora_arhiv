package logger

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Middleware(lg Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := uuid.New().String()
			ctx := r.Context()

			ctx = WithLogger(ctx, lg.WithField(TraceID, requestID))
			ctx = context.WithValue(ctx, CtxTraceID, requestID)
			ctx = metadata.AppendToOutgoingContext(ctx, TraceID, requestID)

			w.Header().Add(TraceID, requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
		resp interface{}, err error) {
		//
		md, ok := metadata.FromIncomingContext(ctx)
		if ok && md.Get(TraceID) != nil {
			lg := Get(ctx).WithField(TraceID, md.Get(TraceID)[0])
			ctx = WithLogger(ctx, lg)
			ctx = context.WithValue(ctx, CtxTraceID, md.Get(TraceID)[0])
		} else {
			requestID := uuid.New().String()
			lg := Get(ctx).WithField(TraceID, requestID)
			ctx = WithLogger(ctx, lg)
			ctx = context.WithValue(ctx, CtxTraceID, requestID)
		}
		return handler(ctx, req)
	}
}
