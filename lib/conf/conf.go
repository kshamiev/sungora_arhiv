package conf

import (
	"flag"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

func Get(cfg interface{}, envPrefix string) error {
	filePath := flag.String("c", "etc/config.dev.yml", "Path to configuration file")
	flag.Parse()
	return GetFile(cfg, *filePath, envPrefix)
}

func GetFile(cfg interface{}, filePath, envPrefix string) error {

	//data, err := os.ReadFile(filePath)
	//if err != nil {
	//	return err
	//}
	//err = yaml.Unmarshal(data, cfg)
	//if err != nil {
	//	return err
	//}

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

	return vip.Unmarshal(cfg)
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
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(cfg, v.Interface(), append(parts, tv)...)
		default:
			envVar := strings.Join(append(parts, tv), ".")
			if err := cfg.BindEnv(envVar); err != nil {
				panic(err)
			}
		}
	}
}
