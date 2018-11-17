package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("login.html"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Home Page</h1>")
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	fmt.Println(username)
	fmt.Println(password)

	tpl.Execute(w, nil)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	loginRouter := r.PathPrefix("/login").Subrouter()
	loginRouter.HandleFunc("", loginGetHandler).Methods("GET")
	loginRouter.HandleFunc("", loginPostHandler).Methods("POST")
	srv := &http.Server{
		Addr:         "127.0.0.1:9000",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
