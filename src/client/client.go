package client

import (
	"sungora/lib/errs"
	"sungora/lib/web"
	"sungora/services/pbsun"
)

var sunClient pbsun.SunClient

func InitSunClient(cfg *web.GRPCConfig) (*web.GRPCClient, error) {
	grpcClient, err := web.NewGRPCClient(cfg)
	if err != nil {
		return nil, errs.NewBadRequest(err)
	}
	sunClient = pbsun.NewSunClient(grpcClient.Conn)
	return grpcClient, nil
}

func GistSunGRPC() pbsun.SunClient {
	return sunClient
}
