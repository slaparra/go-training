package main

import (
	"fmt"
	"github.com/slaparra/go-training/src/08-external-packages/rabbitmq/lib"
)

func main() {

	conn, ch := lib.Connect("amqp://admin:1234@localhost:5678/")

	lib.CreateExchange(ch, lib.ExchangeName())

	for i := 0; i < 5000000; i++ {
		messageType := "command"
		if (i % 2 == 0) {
			messageType = "event"
		}
		routingKey := fmt.Sprintf("a.routing.key.%s.message_name", messageType)
		lib.PublishMessage(
			ch,
			fmt.Sprintf("%s: Hello World: %d", routingKey, i),
			lib.ExchangeName(),
			routingKey,
		)
	}

	defer conn.Close()
	defer ch.Close()
}

