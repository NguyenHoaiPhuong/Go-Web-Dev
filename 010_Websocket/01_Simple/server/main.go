package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage)
	router.HandleFunc("/ws", wsEndpoint)

	log.Println("Server started on port 9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
