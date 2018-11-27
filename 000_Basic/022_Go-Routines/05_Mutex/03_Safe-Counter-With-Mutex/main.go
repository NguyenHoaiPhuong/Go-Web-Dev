package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// SafeCounter uses mutex mux to protect var v
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key
func (sc *SafeCounter) Inc(key string) {
	sc.mux.Lock()
	sc.v[key]++
	sc.mux.Unlock()
	//wg.Done()
}

// Value returns the current value of the counter for the given key
func (sc *SafeCounter) Value(key string) int {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	return sc.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go c.Inc("Akagi")
	}

	//time.Sleep(10 * time.Millisecond)
	fmt.Println(c.Value("Akagi"))
	wg.Wait()
}
