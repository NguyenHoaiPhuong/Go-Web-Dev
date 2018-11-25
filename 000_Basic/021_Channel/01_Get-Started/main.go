package main

import (
	"fmt"
	"time"
)

func write(c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Send to channel value %v\n", i)
		c <- i
		fmt.Printf("Channel received value %v\n", i)
	}
	defer close(c)
}

func read(c chan int) {
	for {
		fmt.Printf("Take out from channel value %v\n", <-c)
		time.Sleep(5 * time.Second)
	}
}

func main() {
	c := make(chan int)
	go write(c)
	read(c)
}
