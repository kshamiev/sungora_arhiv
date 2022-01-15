package client

import (
	"sungora/lib/errs"
	"sungora/lib/web"
	"sungora/services/pbsample"
)

var sampleClient pbsample.SampleClient

func InitSampleClient(cfg *web.GRPCConfig) (*web.GRPCClient, error) {
	grpcClient, err := web.NewGRPCClient(cfg)
	if err != nil {
		return nil, errs.NewBadRequest(err)
	}
	sampleClient = pbsample.NewSampleClient(grpcClient.Conn)
	return grpcClient, nil
}

func GistSampleGRPC() pbsample.SampleClient {
	return sampleClient
}
