package service

import (
	"context"
	"fmt"

	"sungora/types/pbsun"

	"google.golang.org/protobuf/types/known/emptypb"
)

type SunServer struct {
	pbsun.UnimplementedSunServer
}

func (s SunServer) Ping(ctx context.Context, empty *emptypb.Empty) (*pbsun.Test, error) {
	fmt.Println("SunServer.Ping")

	return &pbsun.Test{
		Text: "Funtik",
	}, nil
}

func NewSunServer() pbsun.SunServer {
	return &SunServer{}
}
