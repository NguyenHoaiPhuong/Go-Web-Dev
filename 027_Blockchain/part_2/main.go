package main

import "log"

func print(num int64) {
	bs := IntToHex(num)
	str := string(bs)
	log.Println("bs:", bs)
	log.Println("str:", str)
}

func main() {
	print(1)
	print(10)
	print(100)
	print(1000)
}
