package main

import (
	"fmt"
)

func main() {
	res := func() (i int) {
		defer func() {
			i++
		}()
		return 1
	}()
	fmt.Println(res)
}
