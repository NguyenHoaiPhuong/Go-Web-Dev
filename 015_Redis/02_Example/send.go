package main

import (
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/015_Redis/02_Example/utils"
)

func main() {
	redisCli := utils.RedisClient()
	defer redisCli.Close()

	queueName := "hello_queue"

	for {
		body := utils.ReadFromTerminal()
		intCmd := redisCli.RPush(queueName, body)
		err := intCmd.Err()
		utils.FailOnError(err, "Failed to publish a message")
	}
}
