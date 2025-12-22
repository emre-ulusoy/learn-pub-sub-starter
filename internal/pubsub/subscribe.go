package pubsub

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SubscribeJSON[T any](
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType, // an enum to represent "durable" or "transient"
	handler func(T),
) error {
	ch, _, err := DeclareAndBind(
		conn,
		exchange,
		queueName,
		key,
		queueType,
	)
	if err != nil {
		log.Printf("subscribe error: %v", err)
	}
	incomingMsg, err := ch.Consume(queueName, "", false, false, false, false, nil)

	go func() {
		for msg := range incomingMsg {
			var genericMsg *T
			json.Unmarshal(msg.Body, genericMsg)
			handler(*genericMsg)
			msg.Ack(false)

		}
	}()

	return nil
}
