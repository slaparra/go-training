package lib

import "github.com/streadway/amqp"

func PublishMessage(ch *amqp.Channel, body string, exchange string, routingKey string) {
    err := ch.Publish(
        exchange,           // exchange
        routingKey,         // routing key
        false,    // mandatory
        false,    // immediate
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        })
    FailOnError(err, "Failed to publish a message")
}
