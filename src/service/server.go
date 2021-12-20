package service

import (
	"context"

	"sungora/lib/app"
	"sungora/lib/logger"
	"sungora/lib/response"
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
	lg := logger.Gist(ctx)
	trid := ctx.Value(response.CtxTraceID).(string)
	lg.Info("SunServer.Ping: " + trid)
	lg.Info(s.Span.SpanContext().TraceID.String())
	return &pbsun.Test{
		Text: "Funtik",
	}, nil
}

func NewSunServer() pbsun.SunServer {
	return &SunServer{}
}
