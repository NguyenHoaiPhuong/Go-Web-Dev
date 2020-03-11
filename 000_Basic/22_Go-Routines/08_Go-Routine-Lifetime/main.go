package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	fmt.Println("Start main routine")
	go writeToTempFile()
	time.Sleep(time.Second * 5)
	fmt.Println("Finish main routine")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeToTempFile() {
	fmt.Println("Start writing to temp file")
	defer func() {
		fmt.Println("Finish writing to temp file")
	}()

	i := 0
	for {
		i = i + 1
		line := fmt.Sprintf("Line #%d\n", i)
		err := ioutil.WriteFile("tmp/dat1", []byte(line), 0644)
		check(err)

		fmt.Println(line)
		time.Sleep(time.Second)
	}
}
