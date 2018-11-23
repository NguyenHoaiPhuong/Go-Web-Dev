package main

import (
	"fmt"
)

func main() {
	/*go*/ func() {
		result := doWork1(1, 2)
		fmt.Println("Result 1:", result)
		result = doWork2(result)
		fmt.Println("Result 2:", result)
		result = doWork3(result)
		fmt.Println("Result 3:", result)
	}()

	//time.Sleep(10 * time.Second)
	fmt.Println("hi!")
}

func doWork1(a, b int) int {
	return a + b
}

func doWork2(result int) int {
	return result + 1
}

func doWork3(result int) int {
	return result * 2
}
