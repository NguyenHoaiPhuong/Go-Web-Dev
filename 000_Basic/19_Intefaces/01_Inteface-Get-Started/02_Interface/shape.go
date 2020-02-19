package main

import "fmt"

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Printf("%T\n", s)
}

func totalArea(shapes ...shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
