package main

import (
	"math/big"
	"testing"
)

func TestConvertToFloatByDecimal(t *testing.T) {
	assert := assert.New(t)

	value, isSuccess := new(big.Int).SetString("1000000000000000000", 10)
	assert.Equal(isSuccess, true, "")
	result, _ := ConvertToFloatByDecimal(value, 18).Float64()
	assert.Equal(result, float64(1), "")

	value = new(big.Int).SetInt64(2e18)
	result, _ = ConvertToFloatByDecimal(value, 18).Float64()
	assert.Equal(result, float64(2), "")

	value = new(big.Int).SetInt64(1e18)
	result, _ = ConvertToFloatByDecimal(value, 19).Float64()
	assert.Equal(result, float64(0.1), "")

	value = new(big.Int).SetInt64(1e18)
	result, _ = ConvertToFloatByDecimal(value, 20).Float64()
	assert.Equal(result, float64(0.01), "")

}

func TestConvertToIntByDecimal(t *testing.T) {
	assert := assert.New(t)

	value, isSuccess := new(big.Int).SetString("1000000000000000000", 10)
	assert.Equal(isSuccess, true, "")
	result := ConvertToIntByDecimal(value, 18).Int64()
	assert.Equal(result, int64(1), "")

	value = new(big.Int).SetInt64(1e18)
	result = ConvertToIntByDecimal(value, 17).Int64()
	assert.Equal(result, int64(10), "")

	value = new(big.Int).SetInt64(1e18)
	result = ConvertToIntByDecimal(value, 16).Int64()
	assert.Equal(result, int64(100), "")

}
