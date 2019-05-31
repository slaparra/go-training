package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", onRequestDoSomething)

	http.ListenAndServe(":80", nil)
}

func onRequestDoSomething(res http.ResponseWriter, req *http.Request) {

	sendAMessageToRabbitMq()

	writeHttpResponse(res)
}

func sendAMessageToRabbitMq() {
	conn, err := amqp.Dial("amqp://admin:1234@rabbitmq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello-queue", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare a queue")

	message := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func writeHttpResponse(res http.ResponseWriter) {
	tmpl := template.Must(template.ParseFiles("layout.gohtml"))
	fmt.Println("Message sent")
	data := "Hello world"
	tmpl.Execute(res, data)
}
