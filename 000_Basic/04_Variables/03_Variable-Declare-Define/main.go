package main

import "fmt"

func main() {
	var a string      // Single variable declaration
	a = `Hello World` // Single variable initialization

	var b, c = 2, "Me" // Multiple varaibles declaration and initialization - Infer mixed up data types

	var d, e int = 10, 15 // Multiple varaibles declaration and initialization - Same data type

	f, g := 9.7, false // Shorthand - Infer mixed up data types

	fmt.Printf("a: %v %T\n", a, a)
	fmt.Printf("b: %v %T\n", b, b)
	fmt.Printf("c: %v %T\n", c, c)
	fmt.Printf("d: %v %T\n", d, d)
	fmt.Printf("e: %v %T\n", e, e)
	fmt.Printf("f: %v %T\n", f, f)
	fmt.Printf("g: %v %T\n", g, g)
}
