package app

import (
	"fmt"
	"net"

	"sungora/lib/logger"
	"sungora/lib/request"

	"go.opencensus.io/plugin/ocgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GRPCConfig struct {
	Host         string `yaml:"host"`          // Host
	Port         int    `yaml:"port"`          // Port
	TLS          bool   `yaml:"tls"`           // Использование TLS
	HostOverride string `yaml:"host_override"` // The server name use to verify the hostname returned by TLS handshake
	CaFile       string `yaml:"caFile"`        // The file containing the CA root cert file
	CertFile     string `yaml:"certFile"`      // The TLS cert file
	KeyFile      string `yaml:"keyFile"`       // The TLS key file
}

type GRPCClient struct {
	Conn *grpc.ClientConn
}

// NewGRPCClient создание и старт клиента GRPC
func NewGRPCClient(cfg *GRPCConfig, opts ...grpc.DialOption) (*GRPCClient, error) {
	if cfg.TLS {
		creds, err := credentials.NewClientTLSFromFile(cfg.CaFile, cfg.HostOverride)
		if err != nil {
			return nil, err
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	opts = append(opts, grpc.WithStatsHandler(new(ocgrpc.ClientHandler)))

	var err error
	comp := &GRPCClient{}

	comp.Conn, err = grpc.Dial(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), opts...)
	return comp, err
}

// Close завершение работы клиента GRPC
func (comp *GRPCClient) Close() {
	_ = comp.Conn.Close()
}

type GRPCServer struct {
	Ser       *grpc.Server  //
	Addr      string        // адрес сервера grpc
	chControl chan struct{} // управление ожиданием завершения работы сервера
	lis       net.Listener  //
}

// NewGRPCServer создание и старт сервера GRPC
func NewGRPCServer(cfg *GRPCConfig, opts ...grpc.ServerOption) (*GRPCServer, error) {
	if cfg.TLS {
		creds, err := credentials.NewServerTLSFromFile(cfg.CertFile, cfg.KeyFile)
		if err != nil {
			return nil, err
		}
		opts = append(opts, grpc.Creds(creds))
	}
	opts = append(opts,
		grpc.StatsHandler(new(ocgrpc.ServerHandler)),
		grpc.ChainUnaryInterceptor(logger.Interceptor()),
		grpc.ChainUnaryInterceptor(request.Interceptor()),
	)

	comp := &GRPCServer{
		Ser:       grpc.NewServer(opts...),
		Addr:      fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		chControl: make(chan struct{}),
	}

	var err error
	if comp.lis, err = net.Listen("tcp", comp.Addr); err != nil {
		return nil, err
	}

	go func() {
		_ = comp.Ser.Serve(comp.lis)
		close(comp.chControl)
	}()

	return comp, nil
}

// Close завершение работы сервера GRPC
func (comp *GRPCServer) Close() {
	if comp.lis == nil {
		return
	}

	if err := comp.lis.Close(); err != nil {
		return
	}

	<-comp.chControl
}
