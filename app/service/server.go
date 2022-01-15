package service

import (
	"context"

	"sungora/lib/app"
	"sungora/lib/errs"
	"sungora/lib/logger"
	"sungora/lib/response"
	"sungora/lib/web"
	"sungora/services/pbsample"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SampleServer struct {
	pbsample.UnsafeSampleServer
}

func NewSampleServer(cfg *web.GRPCConfig, opts ...grpc.ServerOption) (*web.GRPCServer, error) {
	grpcServer, err := web.NewGRPCServer(cfg, opts...)
	if err != nil {
		return nil, errs.NewBadRequest(err)
	}
	pbsample.RegisterSampleServer(grpcServer.Ser, &SampleServer{})
	return grpcServer, nil
}

func (ser SampleServer) Ping(ctx context.Context, empty *emptypb.Empty) (*pbsample.Test, error) {
	s := app.NewSpan(ctx)
	s.StringAttribute("description", "qwerty qwerty qwerty")
	defer s.End()
	lg := logger.Gist(ctx)
	trid := ctx.Value(response.CtxTraceID).(string)
	lg.Info("SampleServer.Ping: " + trid)
	lg.Info(s.Span.SpanContext().TraceID.String())
	return &pbsample.Test{
		Text: "Funtik",
	}, nil
}
