package main

import "fmt"

func appendNewMember(mySlice []int) {
	mySlice = append(mySlice, 1, 2, 3)
}

func main() {
	mySlice := []int{10, 20, 30}
	appendNewMember(mySlice)
	fmt.Println(mySlice)
}
