package lib

import (
    "fmt"
    "github.com/streadway/amqp"
    "log"
    "os"
)

func DeclareAndBindQueue(ch *amqp.Channel) (q amqp.Queue) {

    if len(os.Args) < 2 {
        log.Println("Usage: consumer.go [binding_key]...")
        os.Exit(0)
    }

    q, err := ch.QueueDeclare(
        fmt.Sprintf("queue_name_to_bind_%s", os.Args[1]), // name
        true,                                             // durable
        false,                                            // delete when usused
        true,                                             // exclusive
        false,                                            // no-wait
        nil,                                              // arguments
    )
    FailOnError(err, "Failed to declare a queue")

    log.Printf(
        "Binding queue <%s> to exchange <%s> with routing key <%s>",
        q.Name,
        ExchangeName(),
        os.Args[1],
    )

    err = ch.QueueBind(
        q.Name,         // queue name
        os.Args[1],     // routing key
        ExchangeName(), // exchange
        false,
        nil,
    )

    FailOnError(err, "Failed to bind a queue")

    err = ch.Qos(
        20,   // prefetch count
        0,     // prefetch size
        false,      // global
    )

    return
}

func ExchangeName() string {
    return "new_exchange_topic"
}
