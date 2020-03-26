package main

import (
	"fmt"
	"runtime"
)

type student struct {
	name string
}

func main() {
	A := new(student)
	A.name = "Nguyen Hoai Phuong"
	B := A

	fmt.Println("------ Before Removing A -----")

	fmt.Printf("Address of A: %v\n", &A)
	fmt.Printf("Address of B: %v\n", &B)

	fmt.Printf("A point to address: %p\n", A)
	fmt.Printf("B point to address: %p\n", B)

	fmt.Printf("Value of A: %v\n", A)
	fmt.Printf("Value of B: %v\n", B)

	A = nil
	runtime.GC()

	fmt.Println("------ After Removing A -----")

	fmt.Printf("Address of A: %v\n", &A)
	fmt.Printf("Address of B: %v\n", &B)

	fmt.Printf("A point to address: %p\n", A)
	fmt.Printf("B point to address: %p\n", B)

	fmt.Printf("Value of A: %v\n", A)
	fmt.Printf("Value of B: %v\n", B)
}
