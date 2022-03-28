package app

import (
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

const ConfigFilePath = "etc/config.yaml"

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
	data, err := os.ReadFile(dir + "/" + fileConf)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, cfg)
}
