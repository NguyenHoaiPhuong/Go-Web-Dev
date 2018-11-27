package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var counter = 0
var m sync.Mutex

func sayHello(i int) {
	m.Lock()
	fmt.Printf("Hello #%v, counter #%v\n", i, counter)
	counter++
	m.Unlock()

	wg.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go sayHello(i)
	}
	wg.Wait()
}

/* Result:
Hello #1, counter #0
Hello #4, counter #1
Hello #2, counter #2
Hello #3, counter #3
Hello #5, counter #4
Hello #6, counter #5
Hello #7, counter #6
Hello #9, counter #7
Hello #8, counter #8
Hello #0, counter #9
*/
/* Explanation:
A block of codes was locked. Only 1 go routine can access this block, print and change var counter.
*/
