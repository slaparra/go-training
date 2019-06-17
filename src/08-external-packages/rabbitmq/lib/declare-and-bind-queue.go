package lib

import "github.com/streadway/amqp"

func DeclareAndBindQueue(ch *amqp.Channel) (q amqp.Queue) {

    q, err := ch.QueueDeclare(
        "my_new_queue_name_binded_to_new_exchnage_topic",    // name
        true, // durable
        false, // delete when usused
        true,  // exclusive
        false, // no-wait
        nil,   // arguments
    )
    FailOnError(err, "Failed to declare a queue")

    err = ch.QueueBind(
        q.Name,                         // queue name
        "a.routing.key.*",         // routing key
        "new_exchange_topic", // exchange
        false,
        nil,
    )
    FailOnError(err, "Failed to bind a queue")

    err = ch.Qos(
        20,     // prefetch count
        0,     // prefetch size
        false, // global
    )

    return
}
