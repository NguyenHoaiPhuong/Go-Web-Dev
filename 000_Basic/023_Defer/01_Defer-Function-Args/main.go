package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go func() {
		i := 0
		defer fmt.Println(i)
		i++
		wg.Done()
		return
	}()
	wg.Wait()
}
