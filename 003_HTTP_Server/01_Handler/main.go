package main

import (
	"fmt"
	"net/http"
)

type handler_func struct{}

func (h handler_func) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hi. This is handler function")
}

func main() {
	var h handler_func
	http.ListenAndServe(":9000", h)
}
