package service

import (
	"context"

	"sungora/lib/app"
	"sungora/lib/errs"
	"sungora/lib/logger"
	"sungora/lib/web"
	"sungora/services/pbsungora"

	"google.golang.org/grpc"
)

type SungoraServer struct {
	pbsungora.UnsafeSungoraServer
}

func NewSampleServer(cfg *web.GRPCConfig, opts ...grpc.ServerOption) (*web.GRPCServer, error) {
	grpcServer, err := web.NewGRPCServer(cfg, opts...)
	if err != nil {
		return nil, errs.NewBadRequest(err)
	}
	pbsungora.RegisterSungoraServer(grpcServer.Ser, &SungoraServer{})
	return grpcServer, nil
}

func (ser SungoraServer) Ping(ctx context.Context, tt *pbsungora.Test) (*pbsungora.Test, error) {
	s := app.NewSpan(ctx)
	s.StringAttribute("description", "qwerty qwerty qwerty")
	defer s.End()
	lg := logger.Get(ctx)
	lg.Info("SungoraServer.Ping: " + tt.Text)
	return &pbsungora.Test{
		Text: "Funtik",
	}, nil
}
