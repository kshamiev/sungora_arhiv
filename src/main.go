package src

import (
	"flag"
	"fmt"
	"os"

	"sungora/lib/app"
	"sungora/lib/logger"
	"sungora/lib/storage/stpg"
	"sungora/lib/web"
	"sungora/lib/worker"
	"sungora/src/config"
	"sungora/src/service"
)

func Main() {
	flagConfigPath := flag.String("c", "conf/config.yaml", "used for set path to config file")
	flag.Parse()

	// Config загрузка конфигурации & Logger
	lg := logger.NewLogger(nil)
	if err := config.Set(*flagConfigPath); err != nil {
		lg.WithError(err).Fatal("couldn't get config")
	}
	cfg := config.Get()
	lg = logger.Init(&cfg.Lg)

	// Jaeger
	jaeger, err := logger.NewJaeger(&cfg.Jaeger)
	if err != nil {
		lg.WithError(err).Error("jaeger fail")
	}
	defer jaeger.Close()

	// ConnectDB postgres
	if err = stpg.InitConnect(&cfg.Postgresql); err != nil {
		lg.WithError(err).Error("couldn't connect to postgres")
	}

	// Server GRPC
	grpcServer, mux, err := service.NewSampleServer(lg, &cfg.GRPCServer)
	if err != nil {
		lg.WithError(err).Fatal("grpc server error")
	}
	defer grpcServer.Close()
	lg.Info("start grpc server: ", grpcServer.Addr)

	// Client GRPC
	grpcClient, err := service.InitSampleClient(&cfg.GRPCClient)
	if err != nil {
		lg.WithError(err).Fatal("new grpc client error")
	}
	defer grpcClient.Close()

	// Server Web & Handlers
	server, err := web.NewServer(&cfg.ServeHTTP, initRoutes(&cfg.App, mux))
	if err != nil {
		lg.WithError(err).Fatal("new web server error")
	}
	s := fmt.Sprintf("%s://%s:%d", cfg.ServeHTTP.Proto, cfg.ServeHTTP.Host, cfg.ServeHTTP.Port)
	defer server.CloseWait()
	lg.Info("start web server: ", s+"/api/v1/swag/")

	// Workflow
	worker.Init(lg)
	initWorkers()
	defer worker.CloseWait()

	app.Lock(make(chan os.Signal, 1))
}
