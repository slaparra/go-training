# RabbitMq

### Exchange type **Topic** examples

The examples written in this tutorial are for a Topic Exchange routing. Learn more about this kind of configuration in:  
- https://www.rabbitmq.com/tutorials/tutorial-five-go.html

![img](../resources/topic-exchange-routing.jpg)

- [1- Connect to the broker](#connect-to-the-broker)
- [2- Create a topic exchange](#create-an-exchange)
- [3- Produce a message](#publishproducesend-a-message)
- [4- Create a queue](#create-a-queue)
- [5- Bind a queue to the exchange](#bind-a-queue-to-the-exchange)
- [6- Consume messages](#consume-messages)
- [**Practice](#practice)

## Connect to the broker

(Full example in `producer.go` file)

Dial accepts a string in the AMQP URI format and returns a new Connection over TCP using PlainAuth.

```go
conn, err := amqp.Dial("amqp://admin:1234@localhost:5672/")
if err != nil {
    log.Fatalf("%s: %s", "Failed to connect to rabbitMq", err)
}
defer conn.Close()
```

- https://godoc.org/github.com/streadway/amqp#Dial

Next we create a channel, which is where most of the API for getting things done resides:

```go
ch, err := conn.Channel()
if err != nil {
    log.Fatalf("%s: %s", "Failed to open a channel", err)
}
defer ch.Close()
```
- https://godoc.org/github.com/streadway/amqp#Connection.Channel

*It's advisable to use separate connections for
`Channel.Publish` and `Channel.Consume` so not to have TCP pushback on publishing
affect the ability to consume messages.*

## Create an exchange
***ExchangeDeclare** declares an exchange on the server. If the exchange does not already exist, the server will create it. If the exchange exists, the server verifies that it is of the provided type, durability and auto-delete flags.*
- https://godoc.org/github.com/streadway/amqp#Channel.ExchangeDeclare

```go
err = ch.ExchangeDeclare(
        "exchange_topic_name",  // name
        "topic",                // type
        true,                   // durable
        false,                  // auto-deleted
        false,                  // internal
        false,                  // no-wait
        nil,                    // arguments
    )
```

## Publish/produce/send a message
A producer is someone who creates new messages/work. One example is, when a file upload happens, the http handler produces a message for our workers to consume. The http handler/web app is the producer.
```go
err := ch.Publish(
    exchange,           // string
    routingKey,         // string
    false,              // mandatory
    false,              // immediate
    amqp.Publishing{
        ContentType: "text/plain",
        Body:        []byte(body),
    }
)

if err != nil {
    log.Fatalf("%s: %s", "Failed to publish the message", err)
}
```

## Create a queue

(Full example in `consumer.go` file)

***QueueDeclare** declares a queue to hold messages and deliver to consumers.
Declaring creates a queue if it doesn't already exist, or ensures that an
existing queue matches the same parameters.*

- https://godoc.org/github.com/streadway/amqp#Channel.QueueDeclare

```go
q, err := ch.QueueDeclare(
    "my_new_queue_name",    // name
    true,                   // durable
    false,                  // delete when usused
    true,                   // exclusive
    false,                  // no-wait
    nil,                    // arguments
)

if err != nil {
    log.Fatalf("%s: %s", "Failed to declare the queue", err)
}
```

## Bind a queue to the exchange
***QueueBind** binds an exchange to a queue so that publishings to the exchange will
 be routed to the queue when the publishing routing key matches the binding
 routing key.*

- https://godoc.org/github.com/streadway/amqp#Channel.QueueBind
 
```go
//where q is the queue previously created with QueueDeclare
err = ch.QueueBind(
    q.Name,                 // queue name
    "a.routing.key.*",      // routing key
    "new_exchange_topic",   // exchange
    false,
    nil,
)
if err != nil {
    log.Fatalf("%s: %s", "Failed to bind the queue", err)
}
```

*Qos controls how many messages or how many bytes the server will try to keep on
the network for consumers before receiving delivery acks.  The intent of Qos is
to make sure the network buffers stay full between the server and client.*

```go
err = ch.Qos(
    20,         // prefetch count
    0,          // prefetch size
    false,      // global
)
```

## Consume messages
Consuming has a similar meaning to receiving. A consumer is a program that mostly waits to receive messages. Consumers are who consume the messages produced by the producers/publishers. The background workers are consumers.

```go
msgs, err := ch.Consume(
    q.Name,             // queue
    "consumer_name",    // consumer
    true,               // auto ack
    false,              // exclusive
    false,              // no local
    false,              // no wait
    nil,                // args
)
lib.FailOnError(err, "Failed to register a consumer")
```
`Consume` method returns a `chan` where the messages will be received.
```go
func (ch *Channel) Consume(queue, ...) (<-chan Delivery, ...)
```
<small>([More about Channels](09-concurrency.md#channels))</small>

With this feature we can use a goroutine to process the message concurrently: 

```go
go func() {
    for d := range msgs {
        log.Printf(" Consuming: %s", d.Body)
        d.Ack(true)
    }
}()
```
Continues deliveries to the returned **chan (msgs)** Delivery until Channel.Cancel,
Connection.Close, Channel.Close, or an AMQP exception occurs.  Consumers must
range over the chan to ensure all deliveries are received.  Unreceived
deliveries will block all methods on the same connection.


## Practice
- Go to the path: `/src/08-external-packages/rabbitmq`
- Execute `docker-compose up`
- Admin url: `http://localhost:15678`
- Execute: `go run consumer.go "a.routing.key.command.*"`
- Execute: `go run consumer.go "a.routing.key.event.*"`
- Execute: `go run producer.go`


## Related links
- [Producer example](../src/08-external-packages/rabbitmq/producer.go)
- [Consumer example](../src/08-external-packages/rabbitmq/consumer.go)
- [RabbitMq get started](https://www.rabbitmq.com/getstarted.html)
- [Go example with RabbitMq](https://www.rabbitmq.com/tutorials/tutorial-one-go.html)
- [streadway/amqp go package](https://godoc.org/github.com/streadway/amqp)

[(To review)](https://github.com/Pungyeon/go-rabbitmq-example)  
[(To review2)](https://gist.github.com/harrisonturton/c6b62d45e6117d5d03ff44e4e8e1e7f7)