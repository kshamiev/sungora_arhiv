package config

import (
	"sungora/api/miniost"
	"sungora/lib/app"

	"sungora/lib/logger"
	"sungora/lib/storage/stpg"
	"sungora/lib/web"
)

type Config struct {
	App        app.Config           `yaml:"app"`
	Lg         logger.Config        `yaml:"lg"`
	ServeHTTP  web.HttpServerConfig `yaml:"http"`
	Postgresql stpg.Config          `yaml:"psql"`
	Jaeger     logger.JaegerConfig  `yaml:"jaeger"`
	GRPCClient web.GRPCConfig       `yaml:"grpcClient"`
	GRPCServer web.GRPCConfig       `yaml:"grpcServer"`
	Minio      miniost.Config       `yaml:"minio"`
}

var config *Config

func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}

func Init(fileConf string) (*Config, error) {
	cfg := &Config{}
	if err := app.LoadConfig(fileConf, cfg); err != nil {
		return nil, err
	}
	cfg.App.SetDefault()
	config = cfg
	return cfg, nil
}
