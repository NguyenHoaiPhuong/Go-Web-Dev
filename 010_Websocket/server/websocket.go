package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan string)            // broadcast channel

// Upgrader will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  102400,
	WriteBufferSize: 102400,
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// This will determine whether or not an incoming request from a different domain is allowed to connect,
	// and if it isn’t they’ll be hit with a CORS error.
	// For now, we have kept it really simple and simply return true regardless of what host is trying to connect to our endpoint.
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	// defer conn.Close()

	// Register our new client
	clients[conn] = true

	go handleMessages()

	read(conn)
}

func read(conn *websocket.Conn) {
	for {
		// Read message from the browser
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("ReadMessage error: %v\n", err)
			delete(clients, conn)
			conn.Close()
			break
		}

		message := string(msg)

		// print out that message for clarity
		log.Printf("%s sent %s\n", conn.RemoteAddr(), message)

		broadcast <- message
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
