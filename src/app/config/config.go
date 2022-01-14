package config

import (
	"time"

	"sungora/lib/app"
	"sungora/lib/typ"
	"sungora/src/miniost"

	"sungora/lib/logger"
	"sungora/lib/storage/pgsql"
	"sungora/lib/web"
)

const Version = "v1.10.100"

type Config struct {
	App        app.Config           `yaml:"app"`
	Lg         logger.Config        `yaml:"lg"`
	ServeHTTP  web.HttpServerConfig `yaml:"http"`
	Postgresql pgsql.Config         `yaml:"postgresql"`
	Jaeger     logger.JaegerConfig  `yaml:"jaeger"`
	GRPCClient web.GRPCConfig       `yaml:"grpcClient"`
	GRPCServer web.GRPCConfig       `yaml:"grpcServer"`
	Minio      miniost.Config       `json:"minio"`
}

var config *Config

func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}

func Init(fileConf string, cfg *Config) error {
	if err := app.LoadConfig(fileConf, cfg); err != nil {
		return err
	}
	cfg.App.SetDefault()
	cfg.App.Version = Version + " " + time.Now().Format(typ.TimeFormatDMGHIS)
	config = cfg
	return nil
}
