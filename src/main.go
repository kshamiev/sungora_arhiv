package src

import (
	"flag"
	"fmt"
	"log"
	"os"

	"sungora/lib/app"
	"sungora/lib/logger"
	"sungora/lib/storage/pgsql"
	"sungora/lib/web"
	"sungora/lib/worker"
	"sungora/src/config"
)

func Main() {
	flagConfigPath := flag.String("c", "conf/config.yaml", "used for set path to config file")
	flag.Parse()

	// Config загрузка конфигурации & Logger
	cfg := &config.Config{}
	if err := config.Init(*flagConfigPath, cfg); err != nil {
		log.Fatal(err)
	}
	lg := logger.Init(&cfg.Lg)

	// Jaeger
	jaeger, err := logger.NewJaeger(&cfg.Jaeger)
	if err != nil {
		lg.WithError(err).Error("jaeger fail")
	}
	defer jaeger.Close()

	// ConnectDB postgres
	if err = pgsql.InitConnect(&cfg.Postgresql); err != nil {
		lg.WithError(err).Error("couldn't connect to postgres")
	}

	// Server GRPC

	// Client GRPC

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
