package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func pipelineTest() {
	c := gen(2, 3)
	out := square(c)

	fmt.Println(<-out)
	fmt.Println(<-out)
}

func fifo() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := square(in)
	c2 := square(in)

	fmt.Println("Length of c1:", len(c1))
	fmt.Println("Length of c2:", len(c2))

	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}

func main() {
	// pipelineTest()
	// fifo()
	slice := []int{0, 1, 2, 3, 4, 5, 6}
	slice = append(slice[:2], slice[3:]...)
	for _, val := range slice {
		fmt.Println(val)
	}
}
