package conf

import (
	"os"
	"testing"
	"time"

	"sungora/lib/logger/graylog"

	"github.com/stretchr/testify/assert"
)

func TestYaml(t *testing.T) {
	err := os.Setenv("PORTAL_LOG_LEVEL", "trace")
	assert.NoError(t, err, "expected no error")
	cfg := &Config{}
	err = GetDefault(cfg, "config.yml.yml", "portal")
	assert.NoError(t, err, "expected no error")
	assert.Equal(t, cfg.Log.Level, "trace", "expected value from env variable")
}

type Config struct {
	App        App        `yaml:"app"`
	Log        Log        `yaml:"log"`
	ServeHTTP  ServeHTTP  `yaml:"http"`
	Postgresql Postgresql `yaml:"psql"`
}

type App struct {
	Token          string        `yaml:"token"`          //
	SessionTimeout time.Duration `yaml:"sessionTimeout"` //
	Mode           string        `yaml:"mode"`           //
	DirWork        string        `yaml:"dirWork"`        //
	DirWww         string        `yaml:"dirWww"`         //
	Version        string        `yaml:"version"`        //
	SigningKey     string        `yaml:"signingKey"`     //
}

type Log struct {
	Title     string `yaml:"title"`     // title
	Output    string `yaml:"output"`    // enum (stdout | filePathRelative)
	Formatter string `yaml:"formatter"` // enum (json|text)
	Level     string `yaml:"level"`     // enum (error|warning|info|debug|trace)
	IsCaller  bool   `yaml:"is_caller"` // bool
	Hooks     Hooks  `yaml:"hooks"`
}

type Hooks struct {
	Graylog graylog.Config `yaml:"graylog" json:"graylog"`
}

type ServeHTTP struct {
	Proto          string        `yaml:"proto" mapstructure:"proto"`                   // Server Proto
	Host           string        `yaml:"host"`                                         // Server Host
	Port           int           `yaml:"port"`                                         // Server Port
	ReadTimeout    time.Duration `yaml:"readTimeout"`                                  // Время ожидания web запроса в секундах
	WriteTimeout   time.Duration `yaml:"writeTimeout"`                                 // Время ожидания окончания передачи ответа в секундах
	RequestTimeout time.Duration `yaml:"requestTimeout"`                               // Время ожидания окончания выполнения запроса
	IdleTimeout    time.Duration `yaml:"idleTimeout"`                                  // Время ожидания следующего запроса
	MaxHeaderBytes int           `yaml:"maxHeaderBytes" mapstructure:"maxHeaderBytes"` // Максимальный размер заголовка получаемого от браузера клиента в байтах
}

type Postgresql struct {
	Postgres     string   `yaml:"postgres"`
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

func (c *Config) SetDefault() error {
	return nil
}
