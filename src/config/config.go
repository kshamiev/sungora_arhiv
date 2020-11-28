package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"

	"sungora/lib/logger"
	"sungora/lib/storage/stpg"
	"sungora/lib/web"
)

const Version = "v1.0.0"

type Config struct {
	App        App                  `yaml:"app"`
	Lg         logger.Config        `yaml:"lg"`
	ServeHTTP  web.HttpServerConfig `yaml:"http"`
	Postgresql stpg.Config          `yaml:"postgresql"`
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

// загрузка конфигурации приложения
func Set(fileConfig ...string) error {
	config = &Config{}
	for i := range fileConfig {
		if err := ConfigLoad(fileConfig[i], config); err == nil {
			config.App.SetDefault()
			return err
		}
	}
	return nil
}

// ConfigLoad загрузка конфигурации
func ConfigLoad(path string, cfg interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, cfg)
}
