package main

import (
	"fmt"
)

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Visit:")
	visit(numbers, func(i int) {
		fmt.Println(i)
	})

	max := 2
	fmt.Println("Filter numbers bigger than", max)
	xs := filter(numbers, func(n int) bool {
		return n > max
	})
	fmt.Println(xs)
}
