package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Println("x:", x)
	{
		fmt.Println("x:", x)
		y := 77
		fmt.Println("y:", y)
	}
	//fmt.Println("y:", y)	Outside scope of y
}
