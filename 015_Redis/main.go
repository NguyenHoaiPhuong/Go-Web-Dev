package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

const (
	redisServer string = "localhost"
	redisPort   string = "6380"
)

var (
	client *redis.Client
)

func init() {
	// Init redis client
	client = redis.NewClient(&redis.Options{
		Addr:     redisServer + ":" + redisPort,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// ping the redis server
func example1() {
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func example2() {
	err := client.Set("key", "value", 0).Err()
	checkError(err)

	val, err := client.Get("key").Result()
	checkError(err)
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func main() {
	// example1()
	example2()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
