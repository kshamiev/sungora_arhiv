package main

import (
	"fmt"
	"log"
	"os"

	"sample/internal/client"
	"sample/internal/config"
	"sample/internal/handler"
	"sample/internal/service"
	"sample/lib/app"
	"sample/lib/app/sheduler"
	"sample/lib/errs"
	"sample/lib/jaeger"
	"sample/lib/logger"
	"sample/lib/minio"
	"sample/lib/storage/stpg"
)

// @title Sample API
// @description Sample
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
// @tag.description Пример работы с моделью (Пользователи)
// @tag.name Data
// @tag.description Работа с бинарными данными
// @tag.name Websocket
// @tag.description Чат (веб сокет)
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// App загрузка конфигурации & Logger
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(errs.New(err))
	}
	lg := logger.Init(&cfg.Log)

	// ConnectDB postgres
	if err = stpg.InitConnect(&cfg.Postgresql); err != nil {
		lg.Fatal(errs.New(err))
	}

	// Minio
	if err = minio.Init(&cfg.Minio); err != nil {
		lg.Fatal(errs.New(err))
	}

	// Jaeger
	jg, err := jaeger.NewJaeger(&cfg.Jaeger)
	if err != nil {
		lg.Fatal(errs.New(err))
	}
	defer jg.Close()

	// Server GRPC
	var grpcServer *app.GRPCServer
	if grpcServer, err = service.NewSampleServer(&cfg.GRPCServer); err != nil {
		lg.Fatal(err)
	}
	defer grpcServer.Close()
	lg.Info("start grpc server: ", grpcServer.Addr)

	// Client GRPC
	var grpcClient *app.GRPCClient
	if grpcClient, err = client.InitSampleClient(&cfg.GRPCClient); err != nil {
		lg.Fatal(err)
	}
	defer grpcClient.Close()

	// Workflow
	sheduler.Init()
	defer sheduler.CloseWait()

	// Server Web & Handlers
	server, err := app.NewHTTPServer(&cfg.ServeHTTP, handler.Routing(&cfg.App))
	if err != nil {
		lg.WithError(err).Fatal("new web server error")
	}
	defer server.CloseWait()
	lg.Info("start web server: ", fmt.Sprintf(
		"%s://localhost:%d/api/sun/swag/index.html",
		cfg.ServeHTTP.Proto, cfg.ServeHTTP.Port),
	)

	app.Lock(make(chan os.Signal, 1))
}
