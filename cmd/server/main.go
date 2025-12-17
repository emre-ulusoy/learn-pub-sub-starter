package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	pubsub "github.com/bootdotdev/learn-pub-sub-starter/internal"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")

	connStr := "amqp://guest:guest@localhost:5672/"
	conn, err := amqp.Dial(connStr)
	if err != nil {
		log.Fatalf(err.Error())
	}

	defer conn.Close()
	fmt.Println("connection successful")

	newChan, err := conn.Channel()
	if err != nil {
		log.Fatalf(err.Error())
	}

	state := routing.PlayingState{
		IsPaused: true,
	}
	pubsub.PublishJSON[routing.PlayingState](newChan, routing.ExchangePerilDirect, routing.PauseKey, state)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sigChan // INFO: Blocker
	fmt.Println("signal received, shutting down the connection")
	conn.Close()

	fmt.Printf("(Number of goroutines: %d)\n", runtime.NumGoroutine())
}
