package stpg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Postgres     string `yaml:"postgres"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	OcSQLTrace   bool   `yaml:"ocsql_trace"`
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}

func SetConfig(fileConf string) error {
	dir, _ := os.Getwd()
	for {
		data, err := ioutil.ReadFile(dir + "/" + fileConf)
		if err != nil {
			if !strings.Contains(dir, "/") {
				return fmt.Errorf("config '" + fileConf + "' not found")
			}
			dir = filepath.Dir(dir)
		} else {
			var cfg = struct {
				Postgresql Config `json:"postgresql"`
			}{}
			if err := yaml.Unmarshal(data, &cfg); err != nil {
				log.Fatal(err)
			}
			config = &cfg.Postgresql
			return nil
		}
	}
}
