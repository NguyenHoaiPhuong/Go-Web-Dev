package main

import (
	"fmt"
)

func main() {
	for i := 100; i < 110; i++ {
		fmt.Printf("%d - %b - %#x - %q\n", i, i, i, i)
	}
}
