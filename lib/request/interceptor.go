package request

import (
	"context"

	"github.com/google/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"sungora/lib/logger"
	"sungora/lib/response"
)

func LoggerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
		resp interface{}, err error) {
		//
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if md.Get(string(response.CtxToken)) != nil {
				ctx = context.WithValue(ctx, response.CtxToken, md.Get(string(response.CtxToken))[0])
			}
			if md.Get(logger.LogTraceID) != nil {
				lg := logger.Gist(ctx).WithField(logger.LogTraceID, md.Get(logger.LogTraceID)[0])
				ctx = logger.WithLogger(ctx, lg)
				ctx = boil.WithDebugWriter(ctx, lg.Writer())
				ctx = context.WithValue(ctx, logger.CtxTraceID, md.Get(logger.LogTraceID)[0])
			}
		} else {
			requestID := uuid.New().String()
			lg := logger.Gist(ctx).WithField(logger.LogTraceID, requestID)
			ctx = logger.WithLogger(ctx, lg)
			ctx = boil.WithDebugWriter(ctx, lg.Writer())
			ctx = context.WithValue(ctx, logger.CtxTraceID, requestID)
		}
		return handler(ctx, req)
	}
}
