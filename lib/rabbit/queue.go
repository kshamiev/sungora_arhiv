package rabbit

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
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

// ////

func (pro *Producer) Queue(queueName string, data ...string) error {
	_, err := instance.channel.QueueDeclare(
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

	if pro.isConfirm {
		if err := instance.channel.Confirm(false); err != nil {
			return errs.NewBadRequest(err)
		}
		confirms := instance.channel.NotifyPublish(make(chan amqp.Confirmation, 1))
		defer func() {
			ln := len(data)
			for 0 < ln {
				if confirmed := <-confirms; confirmed.Ack {
					log.Printf("confirmed delivery with delivery tag: %d", confirmed.DeliveryTag)
				} else {
					log.Printf("failed delivery of delivery tag: %d", confirmed.DeliveryTag)
				}
				ln--
			}
		}()
	}

	for i := range data {
		if err = instance.channel.Publish(
			pro.exchange, // publish to an exchange
			queueName,    // routing to 0 or more queues
			false,        // mandatory
			false,        // immediate
			amqp.Publishing{
				Headers:         amqp.Table{},
				ContentType:     "text/plain",
				ContentEncoding: "",
				Body:            []byte(data[i]),
				DeliveryMode:    amqp.Transient, // 1=non-persistent, 2=persistent
				Priority:        0,              // 0-9
				// a bunch of application/implementation-specific fields
			},
		); err != nil {
			return errs.NewBadRequest(err)
		}
	}
	return nil
}
