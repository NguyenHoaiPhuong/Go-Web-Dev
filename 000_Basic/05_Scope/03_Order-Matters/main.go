package main

import "fmt"

func main() {
	//fmt.Println(x) // Compile Error: var x MUST be declared before being used
	fmt.Println(y)
	//x := 42
}

var y = 42
