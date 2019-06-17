package main

import (
	"fmt"
	"github.com/slaparra/go-training/src/08-external-packages/rabbitmq/lib"
)

func main() {

	conn, ch := lib.Connect("amqp://admin:1234@localhost:5678/")

	lib.CreateExchange(ch, "new_exchange_topic")

	for i := 0; i < 5000000; i++ {
		lib.PublishMessage(
			ch,
			fmt.Sprintf("Hello World: %d", i),
			"new_exchange_topic",
			"a.routing.key.event_name",
		)
	}

	defer conn.Close()
	defer ch.Close()
}

