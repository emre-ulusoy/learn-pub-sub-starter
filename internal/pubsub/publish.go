package pubsub

import (
	"context"
	"encoding/json"
	// "log"

	amqp "github.com/rabbitmq/amqp091-go"
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
