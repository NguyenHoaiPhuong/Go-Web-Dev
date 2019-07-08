package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	network := "tcp"
	port := ":9000"
	conn, err := net.Dial(network, port)
	if err != nil {
		log.Println("Cannot dial to", port, "due to error", err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprint(conn, text+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
