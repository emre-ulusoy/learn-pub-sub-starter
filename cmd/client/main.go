package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
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

	// ---------- INFO let's put the stuff here
	username, err := gamelogic.ClientWelcome()
	if err != nil {
		log.Fatalf("couldn't get user name: %v", err)
	}

	_, q, err := pubsub.DeclareAndBind(conn, routing.ExchangePerilDirect, routing.PauseKey+"."+username, routing.PauseKey, pubsub.SimpleQueueTransient)
	if err != nil {
		log.Fatalf("could not subscribe to pause %v", err)
	}

	fmt.Printf("queue %s declared and bound!\n", q.Name)

	// -----------------

	// state := routing.PlayingState{
	// 	IsPaused: true,
	// }
	// err = pubsub.PublishJSON(ch, routing.ExchangePerilDirect, routing.PauseKey, state)
	// if err != nil {
	// 	log.Printf("could not publish time: %v", err)
	// }
	// fmt.Println("pause message sent!")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sigChan // INFO: Blocker
	fmt.Println("signal received, RabbitMQ connection closed")
	// conn.Close()
}

func handlerPause(gs *gamelogic.GameState) func(routing.PlayingState) {
	return func(routing.PlayingState) {

	}
}
