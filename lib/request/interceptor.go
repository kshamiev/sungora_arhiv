package request

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"sungora/lib/logger"
	"sungora/lib/response"
)

func LoggerInterceptor(lg logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
		resp interface{}, err error) {
		//
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if md.Get(string(response.CtxToken)) != nil {
				ctx = context.WithValue(ctx, response.CtxToken, md.Get(string(response.CtxToken))[0])
			}
			if md.Get(logger.TraceID) != nil {
				ctx = logger.WithLogger(ctx, lg.WithFields(map[string]interface{}{
					logger.TraceID: md.Get(logger.TraceID)[0],
				}))
			}
			if md.Get(logger.TraceAPI) != nil {
				ctx = logger.WithLogger(ctx, lg.WithFields(map[string]interface{}{
					logger.TraceAPI: md.Get(logger.TraceID)[0],
				}))
			}
		} else {
			ctx = logger.WithLogger(ctx, lg.WithFields(map[string]interface{}{
				logger.TraceID: uuid.New().String(),
			}))
		}

		return handler(ctx, req)
	}
}
