package jaeger

import (
	"context"
	"net/http"
	"path"
	"strings"

	"sample/lib/logger"

	"github.com/go-chi/chi"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

func Observation() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return &ochttp.Handler{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()

				// trick to calculate real route pattern for subrouters
				rctx := chi.RouteContext(ctx)
				nc := chi.Context{}
				rctx.Routes.Match(&nc, r.Method, r.RequestURI)
				httpPath := strings.ReplaceAll(path.Join(nc.RoutePatterns...), "/*/", "/")

				span := trace.FromContext(ctx)
				span.AddAttributes(trace.StringAttribute(ochttp.PathAttribute, httpPath))
				ochttp.SetRoute(ctx, httpPath)

				ctx = context.WithValue(ctx, logger.CtxTraceID, span.SpanContext().TraceID.String())

				next.ServeHTTP(w, r.WithContext(ctx))
			}),
			FormatSpanName: func(r *http.Request) string {
				// trick to calculate real route pattern for subrouters
				rctx := chi.RouteContext(r.Context())
				nc := chi.Context{}
				rctx.Routes.Match(&nc, r.Method, r.RequestURI)
				return strings.ReplaceAll(path.Join(nc.RoutePatterns...), "/*/", "/")
			},
		}
	}
}
