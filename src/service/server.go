package service

import (
	"context"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	"sungora/lib/app"
	"sungora/lib/logger"
	"sungora/lib/request"
	"sungora/lib/web"
	"sungora/pb"
)

type SampleServer struct {
	pb.UnimplementedSampleServer
}

func (s SampleServer) GetVersion(ctx context.Context, _ *empty.Empty) (*pb.Version, error) {
	logger.GetLogger(ctx).Info("SampleServer: GetVersion")
	return &pb.Version{Text: "GetVersion"}, nil
}

// @Tags GRPC
// @Summary создание объекта
// @Router /api/v1/gate [post]
// @Param data body typ.ReqSample true "какой-то объект"
// @Success 200 {object} typ.ResSample "успех"
// @Failure 400 {string} string "провал"
// @Security ApiKeyAuth
func (*SampleServer) PostSample(ctx context.Context, req *pb.ReqSample) (*pb.ResSample, error) {
	app.Dumper(req)
	return &pb.ResSample{}, nil
}

// @Tags GRPC
// @Summary получение объекта
// @Router /api/v1/gate/{id} [get]
// @Param id path string true "ИД объекта"
// @Param name query string false "name"
// @Param flag query bool false "flag"
// @Param hobbit query []string false "hobbit" collectionFormat(multi)
// @Success 200 {object} typ.ResSample "объект"
// @Failure 400 {string} string "провал"
// @Security ApiKeyAuth
func (*SampleServer) GetSample(ctx context.Context, req *pb.ReqSample) (*pb.ResSample, error) {
	return &pb.ResSample{
		Id:     req.Id,
		Name:   req.Name,
		Flag:   req.Flag,
		Hobbit: req.Hobbit,
	}, nil
}

// ////

func NewSampleServer(lg logger.Logger, cfg *web.GRPCConfig) (*web.GRPCServer, http.Handler, error) {
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
	}))
	opts := grpc.ChainUnaryInterceptor(
		request.LoggerInterceptor(lg),
	)
	grpcServer, err := web.NewGRPCServer(cfg, opts)
	if err != nil {
		return nil, nil, err
	}

	ser := &SampleServer{}

	pb.RegisterSampleServer(grpcServer.Ser, ser)
	err = pb.RegisterSampleHandlerServer(context.Background(), mux, ser)
	if err != nil {
		return nil, nil, err
	}

	return grpcServer, mux, nil
}
