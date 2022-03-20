package rabbit

import (
	"encoding/json"
	"errors"
	"sungora/lib/errs"

	"github.com/streadway/amqp"
)

type Producer struct {
	name    string
	confirm chan amqp.Confirmation
}

func NewProducerTopic(exchange string) (*Producer, error) {
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
	return &Producer{
		name:    exchange,
		confirm: instance.channel.NotifyPublish(make(chan amqp.Confirmation, 1)),
	}, nil
}

func (pro *Producer) Topic(routeKey string, data interface{}) error {
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

func NewProducerQueue(queueName string) (*Producer, error) {
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
	return &Producer{
		name:    queueName,
		confirm: instance.channel.NotifyPublish(make(chan amqp.Confirmation, 1)),
	}, nil
}

func (pro *Producer) Queue(data interface{}) error {
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
