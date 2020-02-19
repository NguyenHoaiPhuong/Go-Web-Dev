package main

import (
	"fmt"
)

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	newIncrement := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(newIncrement())
}
