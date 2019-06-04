package main

import "fmt"

func changeMap(myMap map[int]string) {
	myMap[0] = "Something"
}

func addNewMember(myMap map[int]string) {
	myMap[100] = "John"
}

func main() {
	myMap := make(map[int]string)
	myMap[0] = "None"
	myMap[1] = "Akagi"
	myMap[2] = "Yushin"
	fmt.Println(myMap)

	changeMap(myMap)
	fmt.Println("After changeMap:", myMap)

	addNewMember(myMap)
	fmt.Println("After addNewMember:", myMap)
}
