package main

type rectangle struct {
	w float64
	h float64
	shape
}

func (r rectangle) area() float64 {
	return r.w * r.h
}
