package rabbit

import (
	"sync"

	"sungora/lib/errs"

	"github.com/streadway/amqp"
)

type Config struct {
	Uri      string `yaml:"uri"`
	Exchange string `json:"exchange"`
}

type Rabbit struct {
	cfg     *Config
	conn    *amqp.Connection
	channel *amqp.Channel
	wg      sync.WaitGroup
}

var instance *Rabbit

func Init(cfg *Config) (err error) {
	instance = &Rabbit{cfg: cfg}
	if instance.conn, err = amqp.Dial(cfg.Uri); err != nil {
		return errs.New(err)
	}
	if instance.channel, err = instance.conn.Channel(); err != nil {
		return errs.New(err)
	}
	return
}

func CloseWait() {
	instance.wg.Wait()
	_ = instance.channel.Close()
	_ = instance.conn.Close()
}
