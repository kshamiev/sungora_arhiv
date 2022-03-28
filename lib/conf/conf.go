package conf

import (
	"flag"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func Get(cfg interface{}, envPrefix string) {
	filePath := flag.String("c", "etc/config.dev.yml", "Path to configuration file")
	flag.Parse()
	GetFile(cfg, *filePath, envPrefix)
}

func GetFile(cfg interface{}, filePath, envPrefix string) {

	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		panic(err)
	}

	vip := viper.New()
	vip.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if envPrefix != "" {
		vip.SetEnvPrefix(envPrefix)
	}
	bindEnvs(vip, cfg)
	if filePath != "" {
		vip.SetConfigFile(filePath)
		if err := vip.ReadInConfig(); err != nil {
			panic(err)
		}
	}
	if err := vip.Unmarshal(cfg); err != nil {
		panic(err)
	}
}

func bindEnvs(cfg *viper.Viper, cfgStruct interface{}, parts ...string) {
	ifv := reflect.ValueOf(cfgStruct)
	if ifv.Kind() == reflect.Ptr {
		ifv = ifv.Elem()
	}
	for i := 0; i < ifv.NumField(); i++ {
		v := ifv.Field(i)
		t := ifv.Type().Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			if tv, ok = t.Tag.Lookup("yaml"); !ok {
				if tv, ok = t.Tag.Lookup("json"); !ok {
					continue
				}
			}
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(cfg, v.Interface(), append(parts, tv)...)
		default:
			envVar := strings.Join(append(parts, tv), ".")
			err := cfg.BindEnv(envVar)
			if err != nil {
				panic(err)
			}
		}
	}
}
