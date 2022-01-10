package config

import (
	"os"
	"strings"
	"time"
)

// App основная общая конфигурация
type App struct {
	Token          string        `json:"token"`          //
	SessionTimeout time.Duration `yaml:"sessionTimeout"` //
	Domain         string        `yaml:"domain"`         //
	Mode           string        `yaml:"mode"`           //
	DirWork        string        `yaml:"dirWork"`        //
	DirStatic      string        `yaml:"dirStatic"`      //
	DirWww         string        `yaml:"dirWww"`         //
	Version        string        `json:"version"`        //
	SigningKey     string        `yaml:"signingKey"`     //
}

// SetDefault инициализация значениями по умолчанию
func (cfg *App) SetDefault() {
	if cfg == nil {
		cfg = &App{}
	}

	if cfg.Domain == "" {
		cfg.Domain = "localhost"
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
	cfg.DirStatic = cfg.DirWork + cfg.DirStatic
	cfg.DirWww = cfg.DirWork + cfg.DirWww

	// сессия
	if cfg.Token == "" {
		cfg.Token = cfg.Domain
	}
	if cfg.SessionTimeout == 0 {
		cfg.SessionTimeout = time.Duration(14400) * time.Second
	}

	cfg.Version = Version + " " + time.Now().Format(time.RFC3339)
}
