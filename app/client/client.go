package client

import (
	"sungora/lib/errs"
	"sungora/lib/web"
	"sungora/services/pbsungora"
)

var sungoraClient pbsungora.SungoraClient

func InitSungoraClient(cfg *web.GRPCConfig) (*web.GRPCClient, error) {
	grpcClient, err := web.NewGRPCClient(cfg)
	if err != nil {
		return nil, errs.NewBadRequest(err)
	}
	sungoraClient = pbsungora.NewSungoraClient(grpcClient.Conn)
	return grpcClient, nil
}

func GistSungoraGRPC() pbsungora.SungoraClient {
	return sungoraClient
}
