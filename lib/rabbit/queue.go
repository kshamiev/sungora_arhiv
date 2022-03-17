package rabbit

import (
	"github.com/streadway/amqp"
	"log"
	"sungora/lib/errs"
)

func (con *Consumer) Queue(queueName string, h ConsumerHandler) error {
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
	deliveries, err := instance.channel.Consume(
		queueName, // name
		con.Tag,   // consumerTag,
		false,     // noAck
		false,     // exclusive
		false,     // noLocal
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		return errs.NewBadRequest(err)
	}
	instance.wg.Add(1)
	go con.handle(deliveries, h)
	return nil
}

func (con *Consumer) Cancel() error {
	if err := instance.channel.Cancel(con.Tag, true); err != nil {
		return errs.NewBadRequest(err)
	}
	return nil
}

func (con *Consumer) handle(deliveries <-chan amqp.Delivery, h ConsumerHandler) {
	for d := range deliveries {
		_ = d.Ack(false)
		h.Handler(d.Body)
	}
	instance.wg.Done()
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

	if pro.IsConfirm {
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
			pro.Exchange, // publish to an exchange
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
