package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	queue := NewQueue("amqp://rabbitmq:rabbitmq@localhost:5672/", "hello")
	defer queue.Close()

	queue.Consume(func(i string) {
		log.Printf("Received message with second consumer: %s", i)
	})

	queue.Consume(func(i string) {
		log.Printf("Received message with first consumer: %s", i)
	})

	for i := 0; i < 100; i++ {
		log.Println("Sending message...")
		queue.Send(fmt.Sprint("dupa", i))
		time.Sleep(500 * time.Millisecond)
	}
}
