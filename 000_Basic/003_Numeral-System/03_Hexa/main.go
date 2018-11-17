package main

import "fmt"

func main() {
	fmt.Printf("%d - %b - %x \n", 1234, 1234, 1234)
	fmt.Printf("%d - %b - %#x \n", 1234, 1234, 1234)
	fmt.Printf("%d - %b - %#X \n", 1234, 1234, 1234)
	//fmt.Printf("%d \t %b \t %#X \n", 1234, 1234, 1234)
}
