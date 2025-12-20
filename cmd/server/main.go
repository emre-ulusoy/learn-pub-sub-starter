package main

import (
	"fmt"
	"log"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
)

func main() {
	fmt.Println("Starting Peril server...")

	connStr := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(connStr)
	if err != nil {
		log.Fatalf(err.Error())
	}

	defer conn.Close()
	fmt.Println("connection successful (server)")

	publishChan, err := conn.Channel()
	if err != nil {
		log.Fatalf(err.Error())
	}

	state := routing.PlayingState{
		IsPaused: true,
	}
	err = pubsub.PublishJSON(publishChan, routing.ExchangePerilDirect, routing.PauseKey, state)
	if err != nil {
		log.Printf("could not publish time: %v", err)
	}
	fmt.Println("pause message sent!")

	// sigChan := make(chan os.Signal, 1)
	// signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	// <-sigChan // INFO: Blocker
	// fmt.Println("signal received, shutting down the connection")
	// conn.Close()
}
