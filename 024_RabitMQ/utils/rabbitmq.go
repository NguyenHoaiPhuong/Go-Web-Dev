package utils

import (
	"fmt"

	"github.com/streadway/amqp"
)

const (
	// RabbitMqURL : url
	RabbitMqURL string = "localhost:5672"
	// RabbitMQUser user
	RabbitMQUser string = "rabbitmq"
	// RabbitMQPwd password
	RabbitMQPwd string = "rabbitmq"
)

// Dial func
func Dial() *amqp.Connection {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/", RabbitMQUser, RabbitMQPwd, RabbitMqURL))
	FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

// Channel func
func Channel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")
	return ch
}
