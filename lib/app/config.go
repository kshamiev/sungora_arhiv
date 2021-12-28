package app

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func LoadConfig(fileConf string, cfg interface{}) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	for {
		data, err := ioutil.ReadFile(dir + "/" + fileConf)
		if err == nil {
			return yaml.Unmarshal(data, cfg)
		}
		if dir == "/" {
			return fmt.Errorf("config '" + fileConf + "' not found")
		}
		dir = filepath.Dir(dir)
	}
}
