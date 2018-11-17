package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Home Page</h1>")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("login.html"))
	tpl.Execute(w, nil)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	loginRouter := r.PathPrefix("/login").Subrouter()
	loginRouter.HandleFunc("", loginHandler).Methods("GET")
	srv := &http.Server{
		Addr:         "127.0.0.1:9000",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
