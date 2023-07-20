package bluerabbit

import (
	"encoding/json"
	"fmt"

	"semay.com/utils"
)

type sample_email struct {
	Emails  []string `json:"emails"`
	Subject string   `json:"subject"`
	Message string   `json:"message"`
}

func BlueConsumer() {
	// getting app connection and channel
	connection, channel := BrokerConnect()
	defer connection.Close()
	defer channel.Close()

	// opening a channel over the connection established to interact with RabbitMQ
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}

	// declaring consumer with its properties over channel opened
	msgs, err := channel.Consume(
		"blueadmin", // queue
		"",          // consumer
		true,        // auto ack
		false,       // exclusive
		false,       // no local
		false,       // no wait
		nil,         //args
	)
	if err != nil {
		panic(err)
	}

	// process received messages based on their types
	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			switch msg.Type {
			case "BULK_MAIL":
				var email_msg sample_email
				json.Unmarshal(msg.Body, &email_msg)
				utils.SendEmailConsumer(email_msg.Message, email_msg.Subject, email_msg.Emails)
				// utils.SendEmailConsumer(email_msg.Message)
			default:
				fmt.Println("Unknown Task Type")
			}
		}
	}()

	fmt.Println("Waiting for messages...")
	<-forever
}
