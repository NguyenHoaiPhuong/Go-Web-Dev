package main

import (
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/024_RabitMQ/01_Example/utils"
	"github.com/streadway/amqp"
)

func main() {
	conn := utils.Dial()
	defer conn.Close()

	ch := utils.Channel(conn)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	for {
		body := utils.ReadFromTerminal()
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		utils.FailOnError(err, "Failed to publish a message")
	}
}
