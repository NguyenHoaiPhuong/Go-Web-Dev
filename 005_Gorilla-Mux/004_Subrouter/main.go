package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Home Page</h1>")
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Product Page</h1>")
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Article Page</h1>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	productRouter := r.PathPrefix("/Product").Subrouter()
	productRouter.HandleFunc("", productHandler)
	articleRouter := r.PathPrefix("/Article").Subrouter()
	articleRouter.HandleFunc("", articleHandler)

	srv := &http.Server{
		Addr:         "127.0.0.1:9000",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
