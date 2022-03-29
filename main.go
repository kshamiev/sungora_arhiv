package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	"sample/internal/chat"
	"sample/internal/client"
	"sample/internal/config"
	"sample/internal/data"
	"sample/internal/general"
	"sample/internal/service"
	"sample/internal/user"
	"sample/lib/app"
	"sample/lib/app/request"
	"sample/lib/app/worker"
	"sample/lib/errs"
	"sample/lib/jaeger"
	"sample/lib/logger"
	"sample/lib/minio"
	"sample/lib/storage/stpg"
	"sample/lib/tpl"
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
	task := tpl.NewTaskTemplateParse(cfg.App.DirWww)
	if err = task.Action(context.Background()); err != nil {
		lg.Fatal(err)
	}
	worker.AddStart(task)
	defer worker.CloseWait()

	// Server Web & Handlers
	server, err := app.NewHTTPServer(&cfg.ServeHTTP, initDomain(&cfg.App))
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

func initDomain(cfg *config.App) *chi.Mux {
	mid := request.NewMid(cfg.Token, cfg.SigningKey)

	router := chi.NewRouter()
	router.Use(mid.Cors().Handler)
	router.Use(middleware.Recoverer)

	// swagger
	router.Get("/api/sun/swag/*", httpSwagger.Handler())

	// static
	router.Handle("/assets/*", http.FileServer(http.Dir(cfg.DirWww)))

	// pprof
	router.Get("/api/sun/debug/pprof/index", pprof.Index)
	router.Get("/api/sun/debug/pprof/cmdline", pprof.Cmdline)
	router.Get("/api/sun/debug/pprof/profile", pprof.Profile)
	router.Get("/api/sun/debug/pprof/symbol", pprof.Symbol)
	router.Get("/api/sun/debug/pprof/trace", pprof.Trace)
	router.Get("/api/sun/debug/pprof/allocs", pprof.Handler("allocs").ServeHTTP)
	router.Get("/api/sun/debug/pprof/heap", pprof.Handler("heap").ServeHTTP)
	router.Get("/api/sun/debug/pprof/goroutine", pprof.Handler("goroutine").ServeHTTP)
	router.Get("/api/sun/debug/pprof/threadcreate", pprof.Handler("threadcreate").ServeHTTP)
	router.Get("/api/sun/debug/pprof/block", pprof.Handler("block").ServeHTTP)
	router.Get("/api/sun/debug/pprof/mutex", pprof.Handler("mutex").ServeHTTP)

	// domains
	router.Group(func(router chi.Router) {
		router.Use(jaeger.Observation())
		router.Use(logger.Middleware())
		chat.InitDomain(router)
		data.InitDomain(router)
		general.InitDomain(router)
		user.InitDomain(router)
	})

	return router
}
