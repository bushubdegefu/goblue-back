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

// func BlueConsumer() {
// 	// getting app connection and channel
// 	connection, channel := BrokerConnect()
// 	defer connection.Close()
// 	defer channel.Close()

// 	// opening a channel over the connection established to interact with RabbitMQ
// 	channel, err := connection.Channel()
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// declaring consumer with its properties over channel opened
// 	msgs, err := channel.Consume(
// 		"blueadmin", // queue
// 		"",          // consumer
// 		true,        // auto ack
// 		false,       // exclusive
// 		false,       // no local
// 		false,       // no wait
// 		nil,         //args
// 	)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// process received messages based on their types
// 	// for update
// 	forever := make(chan bool)
// 	//  go routine with infinite loop to consume tasks set on rabbit mq
// 	go func() {
// 		for msg := range msgs {
// 			switch msg.Type {
// 			case "BULK_MAIL":
// 				var email_msg sample_email
// 				json.Unmarshal(msg.Body, &email_msg)
// 				utils.SendEmailConsumer(email_msg.Message, email_msg.Subject, email_msg.Emails)
// 				// utils.SendEmailConsumer(email_msg.Message)
// 			default:
// 				fmt.Println("Unknown Task Type")
// 			}
// 		}
// 	}()

// 	fmt.Println("Waiting for messages...")
// 	<-forever
// }

func BlueConsumer() {
    // Getting app connection and channel
    connection, channel, err := BrokerConnect()
    if err != nil {
        fmt.Println("Failed to establish connection:", err)
        return
    }
    defer connection.Close()
    defer channel.Close()

    // Declaring consumer with its properties over the channel opened
    msgs, err := channel.Consume(
        "blueadmin", // queue
        "",          // consumer
        true,        // auto ack
        false,       // exclusive
        false,       // no local
        false,       // no wait
        nil,         // args
    )
    if err != nil {
        fmt.Println("Failed to consume messages:", err)
        return
    }

    // Process received messages based on their types
    // Using a goroutine for asynchronous message consumption
    go func() {
        for msg := range msgs {
            switch msg.Type {
            case "BULK_MAIL":
                var emailMsg sample_email
                err := json.Unmarshal(msg.Body, &emailMsg)
                if err != nil {
                    fmt.Println("Failed to unmarshal message:", err)
                    continue
                }
                utils.SendEmailConsumer(emailMsg.Message, emailMsg.Subject, emailMsg.Emails)
                // utils.SendEmailConsumer(emailMsg.Message)
            default:
                fmt.Println("Unknown Task Type")
            }
        }
    }()

    fmt.Println("Waiting for messages...")
    select {}
}