package main

import "fmt"

func f1() {
	fmt.Println("f1")
	f2()
}

func f2() {
	fmt.Println("f2")
	f3()
}

func f3() {
	fmt.Println("f3")
	f4()
}

func f4() {
	fmt.Println("f4")
	panic("Panic @ f4")
	f5()
}

func f5() {
	fmt.Println("f5")
}

func main() {
	f1()
}
