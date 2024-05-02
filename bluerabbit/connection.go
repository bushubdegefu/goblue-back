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
func BrokerConnect() (*amqp.Connection, *amqp.Channel,error) {

	// con_str := config.Config("RABBIT_BROKER_URL_KUBE")
	// con_str := config.Config("RABBIT_BROKER_URL_KUBE_NODE")
	con_str := config.Config("RABBIT_BROKER_URL")
	
	// connection, err := amqp.Dial(config.Config("RABBIT_BROKER_URL"))
	connection, err := amqp.Dial(con_str)
	if err != nil {
		fmt.Printf("connectin to %v failed due to : %v\n", con_str, err)
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
		connection.Close() // Close the connection if queue declaration fails
        channel.Close()    // Close the channel
		fmt.Printf("creating queue to %v failed due to : %v\n", config.Config("RABBIT_BROKER_URL"), err)

	}
	return connection, channel,nil

}
