package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	defer close(c)
	x, y := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("Before %v\n", i)
		c <- x
		fmt.Printf("After %v\n", i)
		x, y = y, (x + y)
		time.Sleep(time.Second)
	}
}
func main() {
	c := make(chan int, 2)
	//go fibonacci(cap(c), c)
	go fibonacci(10, c)
	for i := range c {
		fmt.Printf("Take out from channel %v\n", i)
	}
}
