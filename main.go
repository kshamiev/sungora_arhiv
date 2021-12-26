package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"sungora/lib/app"
	"sungora/lib/errs"
	"sungora/lib/logger"
	"sungora/lib/storage/pgsql"
	"sungora/lib/web"
	"sungora/lib/worker"
	"sungora/src"
	"sungora/src/client"
	"sungora/src/config"
	"sungora/src/service"
	"sungora/types/pbsun"
)

// @title Sungora API
// @description Sungora
// @version 1.0.0
// @contact.name API Support
// @contact.email konstantin@shamiev.ru
// @license.name Sample License
// @termsOfService http://swagger.io/terms/
//
// @host
// @BasePath /
// @schemes http https
//
// @tag.name General
// @tag.description Общие запросы
// @tag.name User
// @tag.description Пользователи
// @tag.name Websocket
// @tag.description Чат (веб сокет)
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
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
	var grpcServer *web.GRPCServer
	if grpcServer, err = web.NewGRPCServer(&cfg.GRPCServer); err != nil {
		lg.Fatal(errs.NewBadRequest(err))
	}
	pbsun.RegisterSunServer(grpcServer.Ser, service.NewSunServer())
	defer grpcServer.Close()
	lg.Info("start grpc server: ", grpcServer.Addr)

	// Client GRPC
	var grpcClient *web.GRPCClient
	if grpcClient, err = client.InitSunClient(&cfg.GRPCClient); err != nil {
		lg.Fatal(err)
	}
	defer grpcClient.Close()

	// Workflow
	worker.Init()
	defer worker.CloseWait()

	// Server Web & Handlers
	server, err := web.NewServer(&cfg.ServeHTTP, src.Init(&cfg.App))
	if err != nil {
		lg.WithError(err).Fatal("new web server error")
	}
	defer server.CloseWait()
	lg.Info("start web server: ", fmt.Sprintf(
		"%s://%s:%d/api/sun/swag/index.html",
		cfg.ServeHTTP.Proto, cfg.ServeHTTP.Host, cfg.ServeHTTP.Port),
	)

	app.Lock(make(chan os.Signal, 1))
}
