package main

type fibo struct {
	val []int64
}

// Fibo stores all fibo number
var Fibo fibo

const size = 10

func initFibo() {
	Fibo.val = append(Fibo.val, 0, 1)
	for i := 2; i < size; i++ {
		Fibo.val = append(Fibo.val, Fibo.val[i-1]+Fibo.val[i-2])
	}
}
