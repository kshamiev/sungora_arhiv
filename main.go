package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"os"

	"sungora/api/chat"
	"sungora/api/data"
	"sungora/api/general"
	"sungora/api/user"
	"sungora/app/client"
	"sungora/app/config"
	"sungora/app/service"
	"sungora/lib/app"
	"sungora/lib/errs"
	"sungora/lib/logger"
	"sungora/lib/minio"
	"sungora/lib/request"
	"sungora/lib/storage/stpg"
	"sungora/lib/tpl"
	"sungora/lib/web"
	"sungora/lib/worker"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
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
	flagConfigPath := flag.String("c", "conf/config.yaml", "used for set path to config file")
	flag.Parse()

	// Config загрузка конфигурации & Logger
	cfg, err := config.Init(*flagConfigPath)
	if err != nil {
		log.Fatal(errs.NewBadRequest(err))
	}
	lg := logger.Init(&cfg.Lg)
	errs.SetTraceAllow(cfg.App.DirWork)

	// ConnectDB postgres
	if err = stpg.InitConnect(&cfg.Postgresql); err != nil {
		lg.Fatal(errs.NewBadRequest(err))
	}

	// Minio
	if err = minio.Init(&cfg.Minio); err != nil {
		lg.Fatal(errs.NewBadRequest(err))
	}

	// Jaeger
	jaeger, err := logger.NewJaeger(&cfg.Jaeger)
	if err != nil {
		lg.Fatal(errs.NewBadRequest(err))
	}
	defer jaeger.Close()

	// Server GRPC
	var grpcServer *web.GRPCServer
	if grpcServer, err = service.NewSampleServer(&cfg.GRPCServer); err != nil {
		lg.Fatal(err)
	}
	defer grpcServer.Close()
	lg.Info("start grpc server: ", grpcServer.Addr)

	// Client GRPC
	var grpcClient *web.GRPCClient
	if grpcClient, err = client.InitSungoraClient(&cfg.GRPCClient); err != nil {
		lg.Fatal(err)
	}
	defer grpcClient.Close()

	// Workflow
	worker.Init()
	task := tpl.NewTaskTemplateParse(cfg.App.DirWww)
	if err = task.Action(context.Background()); err != nil {
		lg.Fatal(err)
	}
	worker.AddStart(task)
	defer worker.CloseWait()

	// Server Web & Handlers
	server, err := web.NewServer(&cfg.ServeHTTP, initDomain(&cfg.App))
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

func initDomain(cfg *app.Config) *chi.Mux {
	mid := request.NewMid(cfg.Token, cfg.SigningKey)

	router := chi.NewRouter()
	router.Use(mid.Cors().Handler)
	router.Use(middleware.Recoverer)

	// swagger
	router.Get("/api/sun/swag/*", httpSwagger.Handler())

	// static
	router.Handle("/assets/*", http.FileServer(http.Dir(cfg.DirWww)))

	// pprof
	router.Get("/api/sun/debug/pprof/trace", func(w http.ResponseWriter, r *http.Request) {
		pprof.Trace(w, r)
	})
	router.Get("/api/sun/debug/pprof/profile", func(w http.ResponseWriter, r *http.Request) {
		pprof.Profile(w, r)
	})
	router.Get("/api/sun/debug/pprof/symbol", func(w http.ResponseWriter, r *http.Request) {
		pprof.Symbol(w, r)
	})
	router.Get("/api/sun/debug/pprof/allocs", pprof.Handler("allocs").ServeHTTP)
	router.Get("/api/sun/debug/pprof/heap", pprof.Handler("heap").ServeHTTP)
	router.Get("/api/sun/debug/pprof/goroutine", pprof.Handler("goroutine").ServeHTTP)

	// domains
	router.Group(func(router chi.Router) {
		router.Use(mid.Logger())
		router.Use(mid.Observation())
		chat.InitDomain(router)
		data.InitDomain(router)
		general.InitDomain(router)
		user.InitDomain(router)
	})

	return router
}
