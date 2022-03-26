package service

import (
	"context"
	"errors"

	"sungora/lib/errs"
	"sungora/lib/jaeger"
	"sungora/lib/logger"
	"sungora/lib/web"
	"sungora/services/pbsungora"

	"google.golang.org/grpc"
)

type SungoraServer struct {
	pbsungora.UnimplementedSungoraServer
}

func NewSampleServer(cfg *web.GRPCConfig, opts ...grpc.ServerOption) (*web.GRPCServer, error) {
	grpcServer, err := web.NewGRPCServer(cfg, opts...)
	if err != nil {
		return nil, errs.New(err)
	}
	pbsungora.RegisterSungoraServer(grpcServer.Ser, &SungoraServer{})
	return grpcServer, nil
}

func (ser *SungoraServer) Ping(ctx context.Context, tt *pbsungora.Test) (*pbsungora.Test, error) {
	s := jaeger.NewSpan(ctx)
	s.StringAttribute("description", "qwerty qwerty qwerty")
	defer s.End()
	lg := logger.Get(ctx)
	lg.Info("SungoraServer.Ping: " + tt.Text)
	err := errors.New("sample error")
	err = errs.New(err, "user message error")
	lg.WithError(err).Error(err.(*errs.Errs).Response())

	return &pbsungora.Test{
		Text: "Funtik",
	}, nil
}
