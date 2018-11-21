package main

import (
	"fmt"
)

type square struct {
	side float64
}

func (sq square) area() float64 {
	return sq.side * sq.side
}

func main() {
	sq := square{side: 10}
	fmt.Println("Area:", sq.area())
}
