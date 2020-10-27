package main

import (
	"fmt"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func sum(a int, b int) {
	defer recovery()
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	done := make(chan bool)
	go divide(a, b, done)
	<-done
}

func divide(a int, b int, done chan bool) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("recovered:", r)
	// 		done <- true
	// 	}
	// }()
	fmt.Printf("%d / %d = %d", a, b, a/b)
	done <- true
}

func main() {
	sum(5, 0)
	fmt.Println("normally returned from main")
}

// Recover works only when it is called from the same goroutine which is panicking.
// It's not possible to recover from a panic that has happened in a different goroutine.
