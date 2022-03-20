package rabbit

import (
	"encoding/json"
	"errors"
	"sungora/lib/errs"

	"github.com/streadway/amqp"
)

type producer struct {
	name    string
	confirm chan amqp.Confirmation
}

// ////

type ProducerTopic interface {
	Topic(routeKey string, data interface{}) error
}

func NewProducerTopic(exchange string) (ProducerTopic, error) {
	if err := instance.channel.ExchangeDeclare(
		exchange, // name
		"topic",  // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // noWait
		nil,      // arguments
	); err != nil {
		return nil, errs.NewBadRequest(err)
	}
	if err := instance.channel.Confirm(false); err != nil {
		return nil, errs.NewBadRequest(err)
	}
	return &producer{
		name:    exchange,
		confirm: instance.channel.NotifyPublish(make(chan amqp.Confirmation, 1)),
	}, nil
}

func (pro *producer) Topic(routeKey string, data interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return errs.NewBadRequest(err)
	}
	if err = instance.channel.Publish(
		pro.name, // publish to an exchange
		routeKey, // routing to 0 or more queues
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json; charset=UTF-8",
			Body:        d,
		},
	); err != nil {
		return errs.NewBadRequest(err)
	}

	if confirmed := <-pro.confirm; confirmed.Ack {
		return nil
	}
	return errs.NewBadRequest(errors.New("failed delivery message topic"))
}

// ////

type ProducerQueue interface {
	Queue(data interface{}) error
}

func NewProducerQueue(queueName string) (ProducerQueue, error) {
	_, err := instance.channel.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return nil, errs.NewBadRequest(err)
	}
	if err := instance.channel.Confirm(false); err != nil {
		return nil, errs.NewBadRequest(err)
	}
	return &producer{
		name:    queueName,
		confirm: instance.channel.NotifyPublish(make(chan amqp.Confirmation, 1)),
	}, nil
}

func (pro *producer) Queue(data interface{}) error {
	d, err := json.Marshal(data)
	if err != nil {
		return errs.NewBadRequest(err)
	}
	if err = instance.channel.Publish(
		"",       // publish to an exchange
		pro.name, // routing to 0 or more queues
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json; charset=UTF-8",
			Body:        d,
		},
	); err != nil {
		return errs.NewBadRequest(err)
	}

	if confirmed := <-pro.confirm; confirmed.Ack {
		return nil
	}
	return errs.NewBadRequest(errors.New("failed delivery message topic"))
}
