package main

import (
	"testing"
)

var fiboSeries = []int64{
	0,
	1,
	1,
	2,
	3,
	5,
	8,
	13,
	21,
	34,
}

func TestFiboStruct(t *testing.T) {
	initFibo()
	for i := 0; i < size; i++ {
		if Fibo.val[i] != fiboSeries[i] {
			t.Errorf("Fibonacci of %d is incorrect. Expected %d, but got %d", i, fiboSeries[i], Fibo.val[i])
		}
	}
}

func Fibonacci() func() int64 {
	var f0, f1 int64 = 0, 1
	return func() int64 {
		f0, f1 = f1, (f0 + f1)
		return f0
	}
}

func TestFibonacciClosure(t *testing.T) {
	fibo := Fibonacci()
	for i := 1; i < size; i++ {
		v := fibo()
		if v != fiboSeries[i] {
			t.Errorf("Fibonacci of %d is incorrect. Expected %d, but got %d", i, fiboSeries[i], v)
		}
	}
}
