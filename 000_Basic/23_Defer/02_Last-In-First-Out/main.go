package main

import (
	"fmt"
)

func main() {
	func() {
		for i := 0; i < 5; i++ {
			defer fmt.Println(i)
		}
	}()
}
