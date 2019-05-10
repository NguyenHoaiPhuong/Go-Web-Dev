package main

import (
	"fmt"
	"time"
)

func send(c, quit chan int) {
	i := 0
	for {
		select {
		case c <- i:
			fmt.Printf("Send %v\n", i)
			i++
		case <-quit:
			fmt.Println("Quit")
			return
		default:
			fmt.Println("...Waiting...")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// Only receive the first 10 data in the channel c,
// then quit
func receive(c, quit chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("Receive %v\n", <-c)
	}
	quit <- 0
}

// TestSendAndReceive : test send and receive functions
func TestSendAndReceive() {
	c := make(chan int)
	quit := make(chan int)
	go receive(c, quit)
	send(c, quit)
}

// TickToc : tick tock
func TickToc() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("toc!!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	// TestSendAndReceive()
	TickToc()
}
