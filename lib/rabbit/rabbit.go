package rabbit

import (
	"github.com/streadway/amqp"
	"sync"
)

type Config struct {
	Uri      string `yaml:"uri"`
	Exchange string `json:"exchange"`
}

type Rabbit struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	wg      sync.WaitGroup
}

var instance *Rabbit

func Init(cfg *Config) (err error) {
	instance = &Rabbit{}
	if instance.conn, err = amqp.Dial(cfg.Uri); err != nil {
		return
	}
	if instance.channel, err = instance.conn.Channel(); err != nil {
		return
	}
	return
}

func CloseWait() {
	instance.wg.Wait()
	_ = instance.channel.Close()
	_ = instance.conn.Close()
}
