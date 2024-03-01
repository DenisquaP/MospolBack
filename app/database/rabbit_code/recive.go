package rabbit

import (
	"fmt"
	"log"
	emailsender "mospol/internal/functions/email_sender"

	"github.com/streadway/amqp"
)

func Recieve() {
	conn, err := amqp.Dial("amqp://rmuser:rmpassword@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		"emails", // queue
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("\nReceived a message: %s\n", d.Body)
			struc, err := Deserialize(d.Body)
			if err != nil {
				log.Fatal(err)
			}
			emailsender.Sender(struc.Reciever, struc.Message)
			fmt.Printf("\nemail sended to %s\n", struc.Reciever)
		}
	}()

	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
