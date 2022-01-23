package graylog

type Config struct {
	DSN       string   `yaml:"dsn" json:"dsn"`
	Host      string   `yaml:"host" json:"host"`
	Blacklist []string `yaml:"blacklist" json:"blacklist"`
}
