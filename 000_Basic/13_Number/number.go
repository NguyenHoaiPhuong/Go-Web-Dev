package main

import (
	"math/big"
)

// ConvertToFloatByDecimal :
func ConvertToFloatByDecimal(_value *big.Int, _decimal uint64) *big.Float {
	x := new(big.Float).SetInt(_value)
	y := new(big.Float).SetInt(toExp(_decimal))
	return x.Quo(x, y)
}

// ConvertToIntByDecimal :
func ConvertToIntByDecimal(_value *big.Int, _decimal uint64) *big.Int {
	return _value.Div(_value, toExp(_decimal))
}

func toExp(decimal uint64) *big.Int {
	return new(big.Int).Exp(new(big.Int).SetUint64(10), new(big.Int).SetUint64(decimal), nil) // 10^decimal
}
