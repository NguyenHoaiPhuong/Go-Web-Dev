package main

import (
	"fmt"
	"time"
)

func f1(quit chan int) {
	fmt.Println("Start f1")

	time.Sleep(time.Second * 1)
	quit <- 1

	fmt.Println("Stop f1")
}

func f2(quit chan int) {
	fmt.Println("Start f2")

	time.Sleep(time.Second * 1)
	quit <- 2

	fmt.Println("Stop f2")
}

func f3(quit chan int) {
	fmt.Println("Start f3")

	time.Sleep(time.Second * 1)
	quit <- 3

	fmt.Println("Stop f3")
}

func main() {
	quit := make(chan int)
	go f1(quit)
	go f2(quit)
	go f3(quit)

	i := 0
	for {
		select {
		case <-quit:
			i++
			if i == 3 {
				return
			}
		}
	}
}
