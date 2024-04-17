package bluerabbit

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"semay.com/config"
)

// creating connection to the rabbit message broker
// returns the connection based on the connection string
// needs to be closed after using by functions using it
// returns connection and channel struct
func BrokerConnect() (*amqp.Connection, *amqp.Channel) {

	connection, err := amqp.Dial(config.Config("RABBIT_BROKER_URL"))
	// connection, err := amqp.Dial(config.Config("RABBIT_BROKER_URL_KUBE"))
	if err != nil {
		fmt.Printf("connectin to %v failed due to : %v\n", config.Config("RABBIT_BROKER_URL"), err)
	}

	// creating a channel to create a queue
	// instance over the connection we have already
	// established.
	channel, err := connection.Channel()
	if err != nil {
		fmt.Printf("connectin to channel failed due to : %v\n", err)
	}
	// With the instance and declare Queues that we can
	// publish and subscribe to.
	_, err = channel.QueueDeclare(
		"blueadmin", // queue name
		true,        // durable
		false,       // auto delete
		false,       // exclusive
		false,       // no wait
		nil,         // arguments
	)
	if err != nil {
		fmt.Printf("creating queue to %v failed due to : %v\n", config.Config("RABBIT_BROKER_URL"), err)

	}
	return connection, channel

}
