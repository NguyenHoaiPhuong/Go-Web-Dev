package main

import "fmt"

type student struct {
	name  string
	class *class
}

type class struct {
	name string
}

func (c *class) display() {
	fmt.Println("Hello")
}

func main() {
	akagi := &student{name: "Akagi"}
	akagi.class.display()

	fmt.Println(akagi.class)
}
