package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 20; i++ {
		fmt.Printf("foo #%v\n", i)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 20; i++ {
		fmt.Printf("bar #%v\n", i)
		time.Sleep(2 * time.Second)
	}
	wg.Done()
}
