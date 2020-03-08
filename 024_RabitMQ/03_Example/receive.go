package main

import (
	"log"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/024_RabitMQ/03_Example/rabbitmq"
)

func main() {
	conf := &rabbitmq.Config{
		Host:     "localhost",
		Port:     "5672",
		User:     "rabbitmq",
		Password: "rabbitmq",
	}
	queues := []rabbitmq.RabbitMQName{rabbitmq.QueueEmailService, rabbitmq.QueueNotificationService}

	rbmq := rabbitmq.New(conf, queues)
	defer rbmq.Close()
	go rbmq.ReconnectReceiver()

	rbmq.ConsumingMessage()

	go log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}
