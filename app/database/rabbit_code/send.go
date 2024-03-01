package rabbit

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Send() (err error) {
	conn, err := amqp.Dial("amqp://rmuser:rmpassword@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"emails", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	reciever := "piskarev.py@yandex.ru"
	mesg := "New comment"
	body := Message{
		Reciever: reciever,
		Message:  mesg,
	}
	bytes, err := body.Serialize()
	if err != nil {
		return
	}
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/octet-stream",
			Body:        bytes,
		},
	)
	FailOnError(err, "Failed to publish a message")
	fmt.Printf(" [x] Sent %s\n", body)

	return
}
