package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home page")
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Product page")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/Product", productHandler)
	http.ListenAndServe(":9000", r)
}
