package rabbit

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"sungora/lib/errs"
)

func (pro *Producer) Queue(queueName string, data []interface{}) error {
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
		d, err := json.Marshal(data[i])
		if err != nil {
			return errs.NewBadRequest(err)
		}
		if err = instance.channel.Publish(
			pro.exchange, // publish to an exchange
			queueName,    // routing to 0 or more queues
			false,        // mandatory
			false,        // immediate
			amqp.Publishing{
				Headers: amqp.Table{},
				//ContentType:     "text/plain",
				ContentType:     "application/json; charset=UTF-8",
				ContentEncoding: "",
				Body:            d,
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

func (pro *Producer) Exchange(routeKey, queueName string, data []interface{}) error {
	if err := instance.channel.ExchangeDeclare(
		pro.exchange, // name
		"direct",     // type TODO develop feature
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return errs.NewBadRequest(err)
	}

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

	if err = instance.channel.QueueBind(
		queueName,    // name of the queue
		routeKey,     // bindingKey
		pro.exchange, // sourceExchange
		false,        // noWait
		nil,          // arguments
	); err != nil {
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
		d, err := json.Marshal(data[i])
		if err != nil {
			return errs.NewBadRequest(err)
		}
		if err = instance.channel.Publish(
			pro.exchange, // publish to an exchange
			routeKey,     // routing to 0 or more queues
			false,        // mandatory
			false,        // immediate
			amqp.Publishing{
				Headers: amqp.Table{},
				//ContentType:     "text/plain",
				ContentType:     "application/json; charset=UTF-8",
				ContentEncoding: "",
				Body:            d,
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
