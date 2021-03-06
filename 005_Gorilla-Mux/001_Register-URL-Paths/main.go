package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Home page</h1>")
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Product page</h1>")
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Article Page</h1>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/Product", productHandler)
	r.HandleFunc("/Article", articleHandler)
	http.ListenAndServe(":9000", r)
}
