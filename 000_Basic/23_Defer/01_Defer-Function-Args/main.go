package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func passByValue(wg *sync.WaitGroup) {
	wg.Add(1)

	i := 0
	defer fmt.Println("Pass by value:", i)
	i++

	wg.Done()
}

func passByReference(wg *sync.WaitGroup) {
	wg.Add(1)

	i := new(int)
	// fmt.Printf always pass by value
	defer fmt.Printf("Pass by reference: %v\n", *i) // At this state, *i = 0
	*i++

	wg.Done()
}

func main() {
	passByValue(&wg)
	passByReference(&wg)
	wg.Wait()
}
