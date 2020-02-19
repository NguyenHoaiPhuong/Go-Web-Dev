package main

import "fmt"

const nRuns = 10000000000

func main() {
	x := 1.234567891011121314
	y := 0.999999999999999999
	z := 1.000000000000000001

	r1 := x + y
	r1 += z
	fmt.Println(r1)
	for i := 0; i < nRuns; i++ {
		r2 := y + x + z
		if r1 != r2 {
			fmt.Println("r1 =", r1, ", r2 =", r2)
			panic("Random error")
		}
	}
}
