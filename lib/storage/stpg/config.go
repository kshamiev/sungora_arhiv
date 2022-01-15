package stpg

type Config struct {
	Postgres     string   `yaml:"psql"`
	User         string   `yaml:"user"`
	Pass         string   `yaml:"pass"`
	Host         string   `yaml:"host"`
	Port         int      `yaml:"port"`
	Dbname       string   `yaml:"dbname"`
	Sslmode      string   `yaml:"sslmode"`
	Blacklist    []string `yaml:"blacklist"`
	MaxIdleConns int      `yaml:"max_idle_conns"`
	MaxOpenConns int      `yaml:"max_open_conns"`
	OcSQLTrace   bool     `yaml:"ocsql_trace"`
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}
