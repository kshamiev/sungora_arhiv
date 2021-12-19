package service

import (
	"sungora/lib/errs"
	"sungora/lib/web"
	"sungora/types/pbsun"
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

func Gist() pbsun.SunClient {
	return sunClient
}
