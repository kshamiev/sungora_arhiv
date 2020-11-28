package logger

import (
	"net/http"

	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

const TraceID = "trace-id"
const TraceAPI = "api"

func Middleware(log Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			ID := uuid.New().String()
			ctx := WithLogger(req.Context(), log.WithFields(map[string]interface{}{
				TraceID: ID,
			}))
			ctx = metadata.AppendToOutgoingContext(ctx, TraceID, ID)
			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}
