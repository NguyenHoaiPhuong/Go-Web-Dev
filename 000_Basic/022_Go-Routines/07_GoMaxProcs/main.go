package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter = 0
var wg = sync.WaitGroup{}

func main() {
	fmt.Printf("CPUs: %v\n", runtime.NumCPU())
	// Print the number of threads that are available
	// By default, it is equal to the number of cores
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	// Change the number of threads
	runtime.GOMAXPROCS(100)
	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			counter++
			fmt.Printf("go routine #%v\n", counter)
			wg.Done()
		}()
	}
	fmt.Printf("Number of go routine: #%v\n", runtime.NumGoroutine())
	wg.Wait()
}
