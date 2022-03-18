package rabbit

import (
	"fmt"
	"github.com/streadway/amqp"
	"sungora/lib/errs"
	"sungora/lib/logger"
)

func (con *Consumer) Queue(queueName string, h ConsumerHandler) error {
	q, err := instance.channel.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return errs.NewBadRequest(err)
	}
	deliveries, err := instance.channel.Consume(
		queueName, // name
		con.tag,   // consumerTag,
		false,     // noAck
		false,     // exclusive
		false,     // noLocal
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return errs.NewBadRequest(err)
	}
	con.cnt = q.Messages
	con.wg.Add(q.Messages)
	instance.wg.Add(1)
	go con.handle(deliveries, h)
	return nil
}

func (con *Consumer) Exchange(routeKey, queueName string, h ConsumerHandler) error {
	if err := instance.channel.ExchangeDeclare(
		con.exchange, // name
		"direct",     // type TODO develop feature
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return errs.NewBadRequest(err)
	}

	q, err := instance.channel.QueueDeclare(
		queueName, // name of the queue
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return errs.NewBadRequest(err)
	}

	if err = instance.channel.QueueBind(
		queueName,    // name of the queue
		routeKey,     // bindingKey
		con.exchange, // sourceExchange
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return errs.NewBadRequest(err)
	}

	deliveries, err := instance.channel.Consume(
		queueName, // name
		con.tag,   // consumerTag,
		false,     // noAck
		false,     // exclusive
		false,     // noLocal
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return errs.NewBadRequest(err)
	}
	con.cnt = q.Messages
	con.wg.Add(q.Messages)
	instance.wg.Add(1)
	go con.handle(deliveries, h)
	return nil
}

func (con *Consumer) Cancel() error {
	con.wg.Wait()
	instance.wg.Done()
	if err := instance.channel.Cancel(con.tag, true); err != nil {
		return errs.NewBadRequest(err)
	}
	return nil
}

func (con *Consumer) handle(deliveries <-chan amqp.Delivery, h ConsumerHandler) {
	defer func() {
		if rvr := recover(); rvr != nil {
			logger.Gist(con.ctx).Error(errs.NewBadRequest(fmt.Errorf("%+v", rvr)))
			for 0 < con.cnt {
				con.wg.Done()
				con.cnt--
			}
		}
	}()
	for d := range deliveries {
		h.Handler(con.ctx, d.Body)
		_ = d.Ack(false)
		con.wg.Done()
		con.cnt--
	}
}
