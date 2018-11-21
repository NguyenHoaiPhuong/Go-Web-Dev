package main

type square struct {
	side float64
	shape
}

func (sq square) area() float64 {
	return sq.side * sq.side
}
