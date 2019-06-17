package lib

import "github.com/streadway/amqp"

func Connect(url string) (*amqp.Connection, *amqp.Channel) {
    conn, err := amqp.Dial(url)
    FailOnError(err, "Failed to connect to RabbitMQ")

    ch, err := conn.Channel()
    FailOnError(err, "Failed to open a channel")

    return conn, ch
}
