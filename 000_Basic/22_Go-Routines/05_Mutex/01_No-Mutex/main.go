package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var counter = 0

func sayHello(i int) {
	fmt.Printf("Hello #%v, counter #%v\n", i, counter)
	counter++

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
Hello #9, counter #0
Hello #2, counter #0
Hello #5, counter #0
Hello #4, counter #0
Hello #7, counter #0
Hello #6, counter #0
Hello #8, counter #0
Hello #0, counter #0
Hello #1, counter #0
Hello #3, counter #2
*/
/* Explanation:
variable counter wasn't increased (line #13) as expected.
Some Go routines access var counter and print its value before its increment in other go routines.
*/
