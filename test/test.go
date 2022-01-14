package test

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"sungora/lib/errs"
	"sungora/lib/storage/pgsql"
	"sungora/src/app/client"
	"sungora/src/app/config"
	"sungora/src/miniost"
)

func GetEnv() (*config.Config, context.Context) {
	pathConfig := os.Getenv("CONF")
	if pathConfig == "" {
		_, currentFile, _, _ := runtime.Caller(0)
		_ = os.Chdir(filepath.Dir(filepath.Dir(currentFile)))
		pathConfig = "conf/config.yaml"
	}

	cfg := &config.Config{}
	if err := config.Init(pathConfig, cfg); err != nil {
		log.Fatal(errs.NewBadRequest(err))
	}

	// ConnectDB postgres
	if err := pgsql.InitConnect(&cfg.Postgresql); err != nil {
		log.Fatal(errs.NewBadRequest(err))
	}

	// Minio
	if err := miniost.Init(&cfg.Minio); err != nil {
		log.Fatal(errs.NewBadRequest(err))
	}

	// Client GRPC
	if _, err := client.InitSampleClient(&cfg.GRPCClient); err != nil {
		log.Fatal(err)
	}

	return config.Get(), context.Background()
}
