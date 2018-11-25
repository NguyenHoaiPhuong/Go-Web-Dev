package main

import (
	"fmt"
)

func main() {
	sq := square{side: 10}
	rect := rectangle{w: 10, h: 5}
	cir := circle{radius: 5}
	fmt.Println("Square Area:", sq.area())
	info(sq)
	fmt.Println("Rectangle Area:", rect.area())
	info(rect)
	fmt.Println("Circle Area:", cir.area())
	info(cir)
	fmt.Println("Total area:", totalArea(sq, rect, cir))
}
