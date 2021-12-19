package src

import (
	"flag"
	"fmt"
	"log"
	"os"

	"sungora/lib/app"
	"sungora/lib/errs"
	"sungora/lib/logger"
	"sungora/lib/request"
	"sungora/lib/storage/pgsql"
	"sungora/lib/web"
	"sungora/lib/worker"
	"sungora/src/config"
	"sungora/src/service"
	"sungora/types/pbsun"

	"google.golang.org/grpc"
)

func Main() {
	flagConfigPath := flag.String("c", "conf/config.yaml", "used for set path to config file")
	flag.Parse()

	// Config загрузка конфигурации & Logger
	cfg := &config.Config{}
	if err := config.Init(*flagConfigPath, cfg); err != nil {
		log.Fatal(errs.NewBadRequest(err))
	}
	lg := logger.Init(&cfg.Lg)

	// Jaeger
	jaeger, err := logger.NewJaeger(&cfg.Jaeger)
	if err != nil {
		lg.Fatal(errs.NewBadRequest(err))
	}
	defer jaeger.Close()

	// ConnectDB postgres
	if err = pgsql.InitConnect(&cfg.Postgresql); err != nil {
		lg.Fatal(errs.NewBadRequest(err))
	}

	// Server GRPC
	opts := grpc.ChainUnaryInterceptor(
		request.LoggerInterceptor(lg),
	)
	var grpcServer *web.GRPCServer
	if grpcServer, err = web.NewGRPCServer(&cfg.GRPCServer, opts); err != nil {
		lg.Fatal(errs.NewBadRequest(err))
	}
	pbsun.RegisterSunServer(grpcServer.Ser, service.NewSunServer())
	defer grpcServer.Close()
	lg.Info("start grpc server: ", grpcServer.Addr)

	// Client GRPC
	var grpcClient *web.GRPCClient
	if grpcClient, err = service.InitSunClient(&cfg.GRPCClient); err != nil {
		lg.Fatal(err)
	}
	defer grpcClient.Close()

	// Server Web & Handlers
	server, err := web.NewServer(&cfg.ServeHTTP, initRoutes(&cfg.App))
	if err != nil {
		lg.WithError(err).Fatal("new web server error")
	}
	s := fmt.Sprintf("%s://%s:%d", cfg.ServeHTTP.Proto, cfg.ServeHTTP.Host, cfg.ServeHTTP.Port)
	defer server.CloseWait()
	lg.Info("start web server: ", s+"/api/v1/swag/")

	// Workflow
	worker.Init()
	initWorkers()
	defer worker.CloseWait()

	app.Lock(make(chan os.Signal, 1))
}
