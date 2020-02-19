package main

import (
	"fmt"
)

func main() {
	//const sample = "Hello world"
	const sample = "日本語"
	fmt.Println("Sample:", sample)

	fmt.Print("Hexa:\n")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Print("\n")

	fmt.Print("Byte:\n")
	b := []byte(sample)
	for i := 0; i < len(b); i++ {
		fmt.Print(b[i], " ")
	}
	fmt.Print("\n")

	fmt.Print("Rune:\n")
	r := []rune(sample)
	for i := 0; i < len(r); i++ {
		fmt.Print(r[i], " ")
	}
	fmt.Print("\n")

	for idx, rVal := range sample {
		fmt.Printf("%#U starts at byte position %d\n", rVal, idx)
	}
	fmt.Print("\n")
}
