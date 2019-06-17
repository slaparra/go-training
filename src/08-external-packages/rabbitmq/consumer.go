package main

import (
    "github.com/slaparra/go-training/src/08-external-packages/rabbitmq/lib"
    "github.com/streadway/amqp"
    "log"
)

func main() {
    conn, ch := lib.Connect("amqp://admin:1234@localhost:5678/")

    q := lib.DeclareAndBindQueue(ch)

    msgs, err := ch.Consume(
        q.Name, // queue
        "consumer_name",     // consumer
        false,   // auto ack
        false,  // exclusive
        false,  // no local
        false,  // no wait
        nil,    // args
    )
    lib.FailOnError(err, "Failed to register a consumer")

    PrintMessages(msgs)

    defer conn.Close()
    defer ch.Close()
}

func PrintMessages(msgs <-chan amqp.Delivery) {
    forever := make(chan bool)
    go func() {
        for d := range msgs {
            log.Printf(" Consuming: %s", d.Body)
            //time.Sleep(100*time.Millisecond)
            d.Ack(true)
        }
    }()
    log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
    <-forever
}
