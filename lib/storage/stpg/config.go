package stpg

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
