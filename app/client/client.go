package client

import (
	"sungora/lib/app"
	"sungora/lib/errs"
	"sungora/services/pbsungora"
)

var sungoraClient pbsungora.SungoraClient

func InitSungoraClient(cfg *app.GRPCConfig) (*app.GRPCClient, error) {
	grpcClient, err := app.NewGRPCClient(cfg)
	if err != nil {
		return nil, errs.New(err)
	}
	sungoraClient = pbsungora.NewSungoraClient(grpcClient.Conn)
	return grpcClient, nil
}

func GistSungoraGRPC() pbsungora.SungoraClient {
	return sungoraClient
}
