package rabbit

import (
	"context"
	"fmt"
	"sungora/lib/errs"
	"sungora/lib/logger"
	"sungora/lib/typ"

	"github.com/streadway/amqp"
)

type consumer struct {
	queue       string
	consumer    string
	isExclusive bool
	isAck       bool
}

type Consumer interface {
	Handler(ctx context.Context, h func(ctx context.Context, data []byte)) error
	Cancel() error
}

func NewConsumerTopic(exchange, queueName string, routeKey []string) (Consumer, error) {
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

	isDurable := true
	isExclusive := false
	if queueName == "" {
		isDurable = false
		isExclusive = true
	}

	q, err := instance.channel.QueueDeclare(
		queueName,   // name of the queue
		isDurable,   // durable
		false,       // delete when unused
		isExclusive, // exclusive
		false,       // noWait
		nil,         // arguments
	)
	if err != nil {
		return nil, errs.NewBadRequest(err)
	}

	for i := range routeKey {
		if err = instance.channel.QueueBind(
			q.Name,      // name of the queue
			routeKey[i], // bindingKey
			exchange,    // sourceExchange
			false,       // noWait
			nil,         // arguments
		); err != nil {
			return nil, errs.NewBadRequest(err)
		}
	}

	return &consumer{
		queue:       q.Name,
		consumer:    typ.UUIDNew().String(),
		isExclusive: isExclusive,
		isAck:       isExclusive,
	}, nil
}

func NewConsumerQueue(queueName string) (Consumer, error) {
	q, err := instance.channel.QueueDeclare(
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

	return &consumer{
		queue:    q.Name,
		consumer: typ.UUIDNew().String(),
	}, nil
}

// ////

func (con *consumer) Handler(ctx context.Context, h func(ctx context.Context, data []byte)) error {
	deliveries, err := instance.channel.Consume(
		con.queue,    // name
		con.consumer, // consumerTag,
		con.isAck,    // noAck
		false,        // exclusive
		false,        // noLocal
		false,        // noWait
		nil,          // arguments
	)
	if err != nil {
		return errs.NewBadRequest(err)
	}

	instance.wg.Add(1)
	go con.handle(ctx, deliveries, h)
	return nil
}

func (con *consumer) Cancel() error {
	if err := instance.channel.Cancel(con.consumer, false); err != nil {
		return errs.NewBadRequest(err)
	}
	instance.wg.Done()
	return nil
}

func (con *consumer) handle(
	ctx context.Context,
	deliveries <-chan amqp.Delivery,
	h func(ctx context.Context, data []byte),
) {
	defer func() {
		if rvr := recover(); rvr != nil {
			logger.Gist(ctx).Error(errs.NewBadRequest(fmt.Errorf("%+v", rvr)))
		}
	}()
	for d := range deliveries {
		h(ctx, d.Body)
		if !con.isAck {
			_ = d.Ack(false)
		}
	}
}
