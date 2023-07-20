package bluerabbit

import (
	"github.com/streadway/amqp"
	"semay.com/config"
)

// creating connection to the rabbit message broker
// returns the connection based on the connection string
// needs to be closed after using by functions using it
// returns connection and channel struct
func BrokerConnect() (*amqp.Connection, *amqp.Channel) {

	connection, err := amqp.Dial(config.Config("RABBIT_BROKER_URL"))
	if err != nil {
		panic(err)
	}

	// creating a channel to create a qeue
	// instance over the connection we have already
	// established.
	channel, err := connection.Channel()
	if err != nil {
		panic(err)
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
		panic(err)
	}
	return connection, channel

}
