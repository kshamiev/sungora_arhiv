package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"

	"sungora/lib/logger"
	"sungora/lib/storage/pgsql"
	"sungora/lib/web"
)

const Version = "v1.10.100"

type Config struct {
	App        App                  `yaml:"app"`
	Lg         logger.Config        `yaml:"lg"`
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
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	for {
		data, err := ioutil.ReadFile(dir + "/" + fileConf)
		if err == nil {
			if err = yaml.Unmarshal(data, cfg); err != nil {
				return err
			}
			cfg.App.SetDefault()
			config = cfg
			return nil
		}
		if !strings.Contains(dir, "/") {
			return fmt.Errorf("config '" + fileConf + "' not found")
		}
		dir = filepath.Dir(dir)
	}
}
