package config

import (
	"sungora/lib/app"

	"sungora/lib/logger"
	"sungora/lib/storage/pgsql"
	"sungora/lib/web"
)

const Version = "v1.10.100"

type Config struct {
	App App           `yaml:"app"`
	Lg  logger.Config `yaml:"lg"`
	ServeHTTP  web.HttpServerConfig `yaml:"http"`
	Postgresql pgsql.Config         `yaml:"postgresql"`
	Jaeger     logger.JaegerConfig  `yaml:"jaeger"`
	GRPCClient web.GRPCConfig       `yaml:"grpcClient"`
	GRPCServer web.GRPCConfig       `yaml:"grpcServer"`
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
	config = cfg
	return nil
}
