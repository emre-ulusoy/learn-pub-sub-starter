package pubsub

import (
	"context"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type SimpleQueueType int

const (
	Durable SimpleQueueType = iota
	Transient
)

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error {
	valJSON, err := json.Marshal(val)
	if err != nil {
		return err
	}

	pub := amqp.Publishing{
		ContentType: "application/json",
		Body:        valJSON,
	}

	ch.PublishWithContext(context.Background(), exchange, key, false, false, pub)

	return nil
}

func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType, // an enum to represent "durable" or "transient"
) (*amqp.Channel, amqp.Queue, error) {
	ch, err := conn.Channel()
	if err != nil {
		log.Println("error creating channel")
		return nil, amqp.Queue{}, err
	}

	// These vars are initialized for Durable
	isDurable := true
	isAutoDelete := false
	isExclusive := false
	if queueType == Transient {
		isDurable = false
		isAutoDelete = true
		isExclusive = true
	}

	queue, err := ch.QueueDeclare(queueName, isDurable, isAutoDelete, isExclusive, false, nil)
	if err != nil {
		log.Println("error queue declare")
		return nil, amqp.Queue{}, err
	}
	err = ch.QueueBind(queueName, key, exchange, false, nil)
	if err != nil {
		log.Println("error queue bind")
		return nil, amqp.Queue{}, nil
	}
	return ch, queue, nil
}

