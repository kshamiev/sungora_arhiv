package config

import (
	"os"
	"strings"
	"time"

	"sungora/lib/app"
	"sungora/lib/jaeger"
	"sungora/lib/minio"
	"sungora/lib/typ"

	"sungora/lib/logger"
	"sungora/lib/storage/stpg"
	"sungora/lib/web"
)

type Config struct {
	App        App                  `yaml:"app"`
	Lg         logger.Config        `yaml:"lg"`
	ServeHTTP  web.HttpServerConfig `yaml:"http"`
	Postgresql stpg.Config          `yaml:"psql"`
	Jaeger     jaeger.JaegerConfig  `yaml:"jaeger"`
	GRPCClient web.GRPCConfig       `yaml:"grpcClient"`
	GRPCServer web.GRPCConfig       `yaml:"grpcServer"`
	Minio      minio.Config         `yaml:"minio"`
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
func (cfg *App) SetDefault() {
	if cfg == nil {
		cfg = &App{}
	}

	// режим работы приложения
	if cfg.Mode == "" {
		cfg.Mode = "dev"
	}

	// пути
	sep := string(os.PathSeparator)
	if cfg.DirWork == "" {
		cfg.DirWork, _ = os.Getwd()
		sl := strings.Split(cfg.DirWork, sep)
		if sl[len(sl)-1] == "bin" {
			sl = sl[:len(sl)-1]
		}
		cfg.DirWork = strings.Join(sl, sep)
	}
	cfg.DirWww = cfg.DirWork + cfg.DirWww

	// сессия
	if cfg.SessionTimeout == 0 {
		cfg.SessionTimeout = time.Duration(14400) * time.Second
	}

	// версия
	cfg.Version = time.Now().Format(typ.TimeFormatDMGHIS)
}
