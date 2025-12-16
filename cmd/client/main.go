package main

import (
	"fmt"
	"log"

	pubsub "github.com/bootdotdev/learn-pub-sub-starter/internal"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/gamelogic"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
)

func main() {
	fmt.Println("Starting Peril client...")

	connStr := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(connStr)
	if err != nil {
		log.Fatalf(err.Error())
	}

	defer conn.Close()
	fmt.Println("connection successful (client)")

	publishChan, err := conn.Channel()
	if err != nil {
		log.Fatalf("could not create channel: %v", err)
	}

	// INFO let's put the stuff here
	username, err := gamelogic.ClientWelcome()
	if err != nil {
		log.Println("idk how but you fucked up entering a username. we're just gonna call you bib")
	}
	if len(username) == 0 {
		username = "bib"
	}

	state := routing.PlayingState{
		IsPaused: true,
	}
	err = pubsub.PublishJSON(publishChan, routing.ExchangePerilDirect, routing.PauseKey, state)
	if err != nil {
		log.Printf("could not publish time: %v", err)
	}
	fmt.Println("pause message sent!")
}
