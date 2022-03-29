package client

import (
	"sample/lib/app"
	"sample/lib/errs"
	"sample/services/pbsample"
)

var sampleClient pbsample.SampleClient

func InitSampleClient(cfg *app.GRPCConfig) (*app.GRPCClient, error) {
	grpcClient, err := app.NewGRPCClient(cfg)
	if err != nil {
		return nil, errs.New(err)
	}
	sampleClient = pbsample.NewSampleClient(grpcClient.Conn)
	return grpcClient, nil
}

func GistSampleGRPC() pbsample.SampleClient {
	return sampleClient
}
