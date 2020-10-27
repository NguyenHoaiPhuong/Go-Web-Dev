package main

import "fmt"

const numArray = 5
const numString = 4

//Package scope array definition
var integerArray [numArray]int

var stringArray [numString]string

func main() {

	integerArray[0] = 10
	integerArray[1] = 20
	integerArray[2] = 30
	integerArray[3] = 40
	integerArray[4] = 50

	fmt.Println("This is the integer Array: ", integerArray)

	stringArray[0] = "first"
	stringArray[1] = "second"
	stringArray[2] = "third"
	stringArray[3] = "fourth"

	fmt.Println("This is the string array: ", stringArray)

}
