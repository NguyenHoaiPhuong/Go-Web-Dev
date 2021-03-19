package main

import (
	"fmt"
	"strings"
)

var Operator string
const (

)

type NumberOperator struct {
	number   int
	operator string
}

NumberOperator(5, +)
NumberOperator(4, +)
9

NumberOperator(3, *)
NumberOperator(6, /)
18
NumberOperator(2, -)
18/2 = 9

9 + 9 = 18
=================================
NumberOperator(1, "")
18 - 1 = 17

var (
	input = "(5 + 4) * 3 / 2 - 1"

	// => strs = []string{5, +, 4, *, 3, /, 2, -, 1}
	// operators []string{+ * / -}
	// numbers []int{5 4 3 2 1}
)

// O(n)

// func isOperator(str string) bool {

// }

func calculate(input string) int {
	strs := strings.Split(input, " ")
	// fmt.Println(strs)
	ans := 0
	operators := make([]string, 0)
	numbers := make([]int, 0)
	for str := range strs {

	}

	return ans
}

func main() {
	ans := calculate(input)
	fmt.Println(ans)
}

