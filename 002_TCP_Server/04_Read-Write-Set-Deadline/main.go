package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	i := 0

	for {
		conn, err := li.Accept()
		fmt.Println(i)
		i++
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(100 * time.Second))
	if err != nil {
		log.Println("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(conn)
	i := 1000
	for scanner.Scan() {
		fmt.Println(i)
		i++
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s", ln)
	}
	defer conn.Close()

	fmt.Println("***CODE GOT HERE***")
}
