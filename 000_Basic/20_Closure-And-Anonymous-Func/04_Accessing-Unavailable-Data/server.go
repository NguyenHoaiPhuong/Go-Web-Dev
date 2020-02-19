package main

import (
	"net/http"
)

func main() {
	db := NewDatabasae("localhost:5432")
	http.HandleFunc("/hello", hello(db))
	http.ListenAndServe(":9000", nil)
}
