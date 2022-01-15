package app

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"sungora/lib/typ"

	"gopkg.in/yaml.v3"
)

const ConfigFilePath = "conf/config.yaml"

func LoadConfig(fileConf string, cfg interface{}) error {
	if fileConf == "" {
		fileConf = os.Getenv("CONF")
		if fileConf == "" {
			fileConf = ConfigFilePath
		}
	}
	_, currentFile, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filepath.Dir(filepath.Dir(currentFile)))
	_ = os.Chdir(dir)
	data, err := ioutil.ReadFile(dir + "/" + fileConf)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, cfg)
}

// Config основная общая конфигурация
type Config struct {
	Token          string        `yaml:"token"`          //
	SessionTimeout time.Duration `yaml:"sessionTimeout"` //
	Mode           string        `yaml:"mode"`           //
	DirWork        string        `yaml:"dirWork"`        //
	DirWww         string        `yaml:"dirWww"`         //
	Version        string        `yaml:"version"`        //
	SigningKey     string        `yaml:"signingKey"`     //
}

// SetDefault инициализация значениями по умолчанию
func (cfg *Config) SetDefault() {
	if cfg == nil {
		cfg = &Config{}
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
