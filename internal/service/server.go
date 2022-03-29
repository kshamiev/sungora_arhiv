package service

import (
	"context"
	"errors"

	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"sample/lib/app"
	"sample/lib/app/request"
	"sample/lib/errs"
	"sample/lib/jaeger"
	"sample/lib/logger"
	"sample/services/pbsample"
)

type SampleServer struct {
	pbsample.UnimplementedSampleServer
}

func NewSampleServer(cfg *app.GRPCConfig) (*app.GRPCServer, error) {
	opts := []grpc.ServerOption{
		grpc.StatsHandler(new(ocgrpc.ServerHandler)),
		grpc.ChainUnaryInterceptor(logger.Interceptor()),
		grpc.ChainUnaryInterceptor(request.Interceptor()),
	}
	grpcServer, err := app.NewGRPCServer(cfg, opts...)
	if err != nil {
		return nil, errs.New(err)
	}
	pbsample.RegisterSampleServer(grpcServer.Ser, &SampleServer{})
	return grpcServer, nil
}

func (ser *SampleServer) Ping(ctx context.Context, tt *pbsample.Test) (*pbsample.Test, error) {
	s := jaeger.NewSpan(ctx)
	s.StringAttribute("description", "qwerty qwerty qwerty")
	defer s.End()
	lg := logger.Get(ctx)
	lg.Info("SampleServer.Ping: " + tt.Text)
	err := errors.New("sample error")
	err = errs.New(err, "user message error")
	lg.WithError(err).Error(err.(*errs.Errs).Response())

	return &pbsample.Test{
		Text: "Funtik",
	}, nil
}

func (ser *SampleServer) Version(context.Context, *emptypb.Empty) (*pbsample.Test, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
