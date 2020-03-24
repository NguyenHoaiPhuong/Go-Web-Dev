package main

import (
	"fmt"
	"math/big"
)

func main() {
	value, _ := new(big.Int).SetString("100000000000000005", 10)
	result, accuracy := ConvertToFloatByDecimal(value, 18).Float64()
	fmt.Println(result)
	fmt.Println(accuracy)
}
