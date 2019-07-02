package main

import "fmt"

func f() {
	fmt.Println("Start f")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover in f")
		}
	}()

	fmt.Println("Start g")
	g()
	fmt.Println("Return g successfully")
}

func g() {
	defer fmt.Println("defer in g")
	panic("Panic in g")
}

func main() {
	f()
}
