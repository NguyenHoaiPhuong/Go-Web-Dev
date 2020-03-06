package main

import (
	"log"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/015_Redis/02_Example/utils"
)

func main() {
	redisCli := utils.RedisClient()
	defer redisCli.Close()

	queueName := "hello_queue"

	forever := make(chan bool)

	go func() {
		for {
			strCmd := redisCli.RPop(queueName)
			value, err := strCmd.Result()
			if err == nil {
				log.Printf("Received : %s", value)
			}
		}

	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
