package main

import (
	"fmt"
	"sync"
)

func main() {
	// mutex := sync.RWMutex{}

	myMap := make(map[int]int)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 50; i++ {
			fmt.Println("Writing", i)
			// mutex.Lock()
			myMap[i] = i + 1000
			// mutex.Unlock()
		}
		wg.Done()
	}()

	// time.Sleep(time.Second * 2)

	for i := 0; i < 50; i++ {
		fmt.Println("Reading", i)
		// mutex.Lock()
		fmt.Println(myMap[i])
		// mutex.Unlock()
	}

	wg.Wait()
}
