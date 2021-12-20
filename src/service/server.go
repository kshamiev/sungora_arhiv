package service

import (
	"context"
	"fmt"

	"sungora/lib/app"
	"sungora/types/pbsun"

	"google.golang.org/protobuf/types/known/emptypb"
)

type SunServer struct {
	pbsun.UnimplementedSunServer
}

func (ser SunServer) Ping(ctx context.Context, empty *emptypb.Empty) (*pbsun.Test, error) {
	s := app.NewSpan(ctx)
	s.StringAttribute("description", "qwerty qwerty qwerty")
	defer s.End()
	fmt.Println("SunServer.Ping")
	return &pbsun.Test{
		Text: "Funtik",
	}, nil
}

func NewSunServer() pbsun.SunServer {
	return &SunServer{}
}
