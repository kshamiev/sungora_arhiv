package logger

import (
	"context"

	"google.golang.org/grpc"
)

// GRPCUnaryServerInterceptor wraps log to unary RPC interceptor on the server
func GRPCUnaryServerInterceptor(log Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log = log.WithField("method", info.FullMethod)

		resp, err := handler(WithLogger(ctx, log), req)
		if err != nil {
			log.WithError(err).Info("grpc completed with error")
		} else {
			log.Info("grpc completed")
		}

		return resp, err
	}
}

// GRPCStreamServerInterceptor wraps log to streaming RPC interceptor on the server
func GRPCStreamServerInterceptor(log Logger) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log = log.WithField("method", info.FullMethod)

		wrapped := WrapServerStream(stream)
		wrapped.WrappedContext = WithLogger(stream.Context(), log)

		err := handler(srv, wrapped)

		if err != nil {
			log.WithError(err).Info("grpc completed with error")
		} else {
			log.Info("grpc completed")
		}

		return err
	}
}

// WrappedServerStream is a thin wrapper around gRPC ServerStream that allows modifying context.
type WrappedServerStream struct {
	grpc.ServerStream
	// WrappedContext is the wrapper's own Context. You can assign it.
	WrappedContext context.Context
}

// Context returns the wrapper's WrappedContext, overwriting the nested gRPC ServerStream.Context()
func (w *WrappedServerStream) Context() context.Context {
	return w.WrappedContext
}

// WrappedServerStream returns a ServerStream that has the ability to overwrite context.
func WrapServerStream(stream grpc.ServerStream) *WrappedServerStream {
	if existing, ok := stream.(*WrappedServerStream); ok {
		return existing
	}

	return &WrappedServerStream{ServerStream: stream, WrappedContext: stream.Context()}
}
