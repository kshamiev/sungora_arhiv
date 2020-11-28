package service

import (
	"google.golang.org/grpc"

	"sungora/lib/web"
	"sungora/pb"
)

var sampleClient pb.SampleClient

func InitSampleClient(cfg *web.GRPCConfig) (*web.GRPCClient, error) {
	grpcClient, err := web.NewGRPCClient(cfg)
	if err != nil {
		return nil, err
	}
	sampleClient = pb.NewSampleClient(grpcClient.Conn)
	return grpcClient, nil
}

func GetSampleClient() pb.SampleClient {
	if sampleClient == nil {
		sampleClient = pb.NewSampleClient(&grpc.ClientConn{})
	}
	return sampleClient
}
