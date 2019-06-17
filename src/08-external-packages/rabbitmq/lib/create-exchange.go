package lib

import (
    "github.com/streadway/amqp"
    "log"
)

func CreateExchange(ch *amqp.Channel, name string) {
    err := ch.ExchangeDeclare(
        name, // name
        "topic",      // type
        true,          // durable
        false,         // auto-deleted
        false,         // internal
        false,         // no-wait
        nil,           // arguments
    )
    FailOnError(err, "Failed to declare an exchange")
}

func FailOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
    }
}
