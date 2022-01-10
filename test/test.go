package test

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"sungora/lib/errs"
	"sungora/src/config"
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

	return config.Get(), context.Background()
}
