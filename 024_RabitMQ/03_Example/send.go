package main

import (
	"fmt"
	"log"
	"time"

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

	go rbmq.ReconnectSender()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	for i := 0; i < 100; i++ {
		log.Println("Sending message...")
		rbmq.SendMessage(fmt.Sprint("dupa", i))
		time.Sleep(500 * time.Millisecond)
	}
}
