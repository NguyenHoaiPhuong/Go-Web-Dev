package main

import (
	"fmt"
	"time"
)

func main() {
	go foo()
	go bar()
	time.Sleep(10 * time.Second)
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("foo:", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("bar:", i)
		time.Sleep(100 * time.Millisecond)
	}
}
