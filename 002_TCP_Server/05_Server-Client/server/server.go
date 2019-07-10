package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	fmt.Println("Launching server ...")

	// listen to port 9000
	port := ":9000"
	network := "tcp"
	li, err := net.Listen(network, port)
	if err != nil {
		log.Println("Cannot listen to", port, "due to error", err)
		return
	}

	conn, err := li.Accept()
	if err != nil {
		log.Println("Cannot accept connection on port", port, "due to error", err)
		return
	}

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Message received:", message)
		newMessage := strings.ToUpper(message)
		conn.Write([]byte(newMessage + "\n"))
	}
}
