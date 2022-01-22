package service

import (
	"context"

	"sungora/lib/app"
	"sungora/lib/errs"
	"sungora/lib/logger"
	"sungora/lib/web"
	"sungora/services/pbsungora"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (ser SungoraServer) Ping(ctx context.Context, empty *emptypb.Empty) (*pbsungora.Test, error) {
	s := app.NewSpan(ctx)
	s.StringAttribute("description", "qwerty qwerty qwerty")
	defer s.End()
	lg := logger.Gist(ctx)
	lg.Info("SungoraServer.Ping")
	return &pbsungora.Test{
		Text: "Funtik",
	}, nil
}
