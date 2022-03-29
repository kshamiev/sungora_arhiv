package config

import (
	"flag"
	"os"
	"strings"
	"time"

	"sample/lib/app"
	"sample/lib/conf"
	"sample/lib/jaeger"
	"sample/lib/logger"
	"sample/lib/minio"
	"sample/lib/storage/stpg"
)

type Config struct {
	App        App                  `yaml:"app"`
	Log        logger.Config        `yaml:"log"`
	ServeHTTP  app.HttpServerConfig `yaml:"http"`
	Postgresql stpg.Config          `yaml:"psql"`
	Jaeger     jaeger.JaegerConfig  `yaml:"jaeger"`
	GRPCClient app.GRPCConfig       `yaml:"grpcClient"`
	GRPCServer app.GRPCConfig       `yaml:"grpcServer"`
	Minio      minio.Config         `yaml:"minio"`
}

var config *Config

func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}

func Init() (*Config, error) {
	filePath := flag.String("c", "config.yml", "Path to configuration file")
	flag.Parse()
	config = &Config{}
	if err := conf.GetDefault(config, *filePath, ""); err != nil {
		return nil, err
	}
	return config, nil
}

// App основная общая конфигурация
type App struct {
	Token          string        `yaml:"token"`          //
	SessionTimeout time.Duration `yaml:"sessionTimeout"` //
	Mode           string        `yaml:"mode"`           //
	DirWork        string        `yaml:"dirWork"`        //
	DirWww         string        `yaml:"dirWww"`         //
	Version        string        `yaml:"version"`        //
	SigningKey     string        `yaml:"signingKey"`     //
}

// SetDefault инициализация значениями по умолчанию
func (cfg *Config) SetDefault() error {
	if cfg == nil {
		cfg = &Config{}
	}

	// режим работы приложения
	if cfg.App.Mode == "" {
		cfg.App.Mode = "dev"
	}

	// пути
	sep := string(os.PathSeparator)
	if cfg.App.DirWork == "" {
		cfg.App.DirWork, _ = os.Getwd()
		sl := strings.Split(cfg.App.DirWork, sep)
		if sl[len(sl)-1] == "bin" {
			sl = sl[:len(sl)-1]
		}
		cfg.App.DirWork = strings.Join(sl, sep)
	}
	cfg.App.DirWww = cfg.App.DirWork + cfg.App.DirWww

	// сессия
	if cfg.App.SessionTimeout == 0 {
		cfg.App.SessionTimeout = time.Duration(14400) * time.Second
	}

	// версия
	cfg.App.Version = time.Now().Format(conf.TimeFormatDMGHIS)
	return nil
}
