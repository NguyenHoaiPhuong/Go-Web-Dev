package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "Hello"

	go func() {
		fmt.Printf("Anonymous func without args: %v\n", msg)
	}()
	go func(msg string) {
		fmt.Printf("Anonymous func with args: %v\n", msg)
	}(msg)

	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}
