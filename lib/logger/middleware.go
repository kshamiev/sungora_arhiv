package logger

import (
	"context"
	"net/http"
	"path"
	"strings"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Middleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			traceID, ok := ctx.Value(CtxTraceID).(string)
			if !ok {
				traceID = uuid.New().String()
				ctx = context.WithValue(ctx, CtxTraceID, traceID)
			}

			rctx := chi.RouteContext(ctx)
			nc := chi.Context{}
			rctx.Routes.Match(&nc, r.Method, r.RequestURI)
			p := strings.ReplaceAll(path.Join(nc.RoutePatterns...), "/*/", "/")

			ctx = WithLogger(ctx, Get(ctx).WithField(TraceID, traceID).WithField(Api, p))
			ctx = metadata.AppendToOutgoingContext(ctx, TraceID, traceID)
			ctx = metadata.AppendToOutgoingContext(ctx, Api, p)

			w.Header().Add(TraceID, traceID)
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
			ctx = context.WithValue(ctx, CtxTraceID, md.Get(TraceID)[0])
			ctx = WithLogger(ctx, Get(ctx).WithField(TraceID, md.Get(TraceID)[0]).WithField(Api, md.Get(Api)[0]))
		} else {
			traceID, ok := ctx.Value(CtxTraceID).(string)
			if !ok {
				traceID = uuid.New().String()
				ctx = context.WithValue(ctx, CtxTraceID, traceID)
			}
			ctx = WithLogger(ctx, Get(ctx).WithField(TraceID, traceID).WithField(Api, info.FullMethod))
			ctx = metadata.AppendToOutgoingContext(ctx, TraceID, traceID)
			ctx = metadata.AppendToOutgoingContext(ctx, Api, info.FullMethod)
		}
		return handler(ctx, req)
	}
}
