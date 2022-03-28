package conf

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapStructure(t *testing.T) {
	err := os.Setenv("PORTAL_LOG_LEVEL", "trace")
	assert.NoError(t, err, "expected no error")
	cfg := &Configuration{}
	err = Get(cfg, "config.mapstructure.yml", "portal")
	assert.NoError(t, err, "expected no error")
	assert.Equal(t, cfg.Lg.Level, "trace", "expected value from env variable")
}

type (
	Configuration struct {
		AuthGrpcServerEndpoint     string        `mapstructure:"AuthGrpcServerEndpoint"`
		AuthGrpcMiddlewareEndpoint string        `mapstructure:"AuthGrpcMiddlewareEndpoint"`
		Database                   Database      `mapstructure:"Database"`
		Auth                       Auth          `mapstructure:"Auth"`
		ClientLogging              ClientLogging `mapstructure:"ClientLogging"`
		Lg                         Lg            `mapstructure:"Log"`
	}
	Database struct {
		Url      string `mapstructure:"Url"`
		Name     string `mapstructure:"Name"`
		Login    string `mapstructure:"Login"`
		Password string `mapstructure:"Password"`
		TLS      string `mapstructure:"TLS"`
		Timeout  int    `mapstructure:"Timeout"`
		RootCert string `mapstructure:"RootCert"`
	}
	Auth struct {
		AuthType           string `mapstructure:"AuthType"`
		AuthURI            string `mapstructure:"AuthURI"`
		RedirectURI        string `mapstructure:"RedirectURI"`
		TokenURI           string `mapstructure:"TokenURI"`
		RefreshURI         string `mapstructure:"RefreshURI"`
		ClientId           string `mapstructure:"ClientId"`
		ClientSecret       string `mapstructure:"ClientSecret"`
		TestUser           string `mapstructure:"TestUser"`
		TestCounterpartyId int32  `mapstructure:"TestCounterpartyId"`
		Superuser          string `mapstructure:"Superuser"`
	}
	ClientLogging struct {
		Enable    bool  `mapstructure:"Enable"`
		QueueSize int64 `mapstructure:"QueueSize"`
	}
	Lg struct {
		Level  string `mapstructure:"Level"`
		Format string `mapstructure:"Format"`
		Title  string `mapstructure:"Title"`
	}
)

func (c *Configuration) SetDefault() error {
	return nil
}
