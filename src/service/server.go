package service

import (
	"context"

	"sungora/types/pbsun"

	"google.golang.org/protobuf/types/known/emptypb"
)

type SunServer struct {
	pbsun.UnimplementedSunServer
}

func (s SunServer) Ping(ctx context.Context, empty *emptypb.Empty) (*pbsun.Test, error) {
	return &pbsun.Test{
		Text: "Funtik",
	}, nil
}

func NewSunServer() pbsun.SunServer {
	return &SunServer{}
}
