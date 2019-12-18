package model

import "strconv"

// Person : struct
type Person struct {
	Name string
	Age  int
}

// Hello : introduce
func Hello(name string, age int) (string, error) {
	return "I am " + name + " !" + "I am " + strconv.Itoa(age) + "years old", nil
}
