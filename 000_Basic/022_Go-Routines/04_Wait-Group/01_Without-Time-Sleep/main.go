package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	msg := "Hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Printf("%v\n", msg)
		wg.Done()
	}(msg)
	wg.Wait()
}
