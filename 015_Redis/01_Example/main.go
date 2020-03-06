package main

import (
	"fmt"
	"time"

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

// Set variable without expiration time
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

// Set variable with expiration time
func example3() {
	cmd := client.SetNX("myKey", "myValue", 10*time.Second)
	fmt.Println(cmd)
	fmt.Println(cmd.Args())
	err := cmd.Err()
	checkError(err)
}

// append a value into an existing variable
func example4() {
	cmd := client.Append("name", "Seo")
	fmt.Println(cmd)
	err := cmd.Err()
	checkError(err)
}

// list
func example5() {
	listName := "task_queue"
	intCmd := client.RPush(listName, "test_1", "test_2", "test_3")
	fmt.Println(intCmd)
	err := intCmd.Err()
	checkError(err)

	intCmd = client.LPush(listName, "test_4", "test_5")
	fmt.Println(intCmd)
	err = intCmd.Err()
	checkError(err)

	strCmd := client.LRange(listName, 0, -1)
	values, err := strCmd.Result()
	checkError(err)
	for _, val := range values {
		fmt.Println(val)
	}
}

// remove variable from redis-server
func example6() {
	listName := "task_queue"
	intCmd := client.Del(listName)
	fmt.Println(intCmd)
	res, err := intCmd.Result()
	checkError(err)
	fmt.Println("res:", res)

	strCmd := client.LRange(listName, 0, -1)
	values, err := strCmd.Result()
	checkError(err)
	for _, val := range values {
		fmt.Println(val)
	}
}

func main() {
	// example1() // ping the server
	// example2() // set variable without expiration time
	// example3() // set variable with expiration time
	// example4() // append new value to an existing variable
	// example5() // list
	example6()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
