package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)

	log.Println("Server starts on port 8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}