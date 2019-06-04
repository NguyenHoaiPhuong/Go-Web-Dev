package main

import "fmt"

func changeSlice(mySlice []int) {
	mySlice[0] = 100
}

func main() {
	mySlice := []int{10, 20, 30}
	changeSlice(mySlice)
	fmt.Println(mySlice)
}
