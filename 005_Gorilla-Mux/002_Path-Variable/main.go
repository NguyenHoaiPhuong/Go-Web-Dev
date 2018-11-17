package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Home Page</h1>")
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Product Page</h1>")
}

func articleCategoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Article Category Page</h1>")

	vars := mux.Vars(r)
	fmt.Fprintln(w, "<p>"+"category:"+vars["category"]+"</p>"+"<br>")
	fmt.Fprintln(w, "<p>"+"sort:"+vars["sort"]+"</p>")
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Article Page</h1>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/Product/{key}", productHandler)
	r.HandleFunc("/Article/{category}/{sort:(?:asc|desc|new)}", articleCategoryHandler)
	r.HandleFunc("/Article/{category}/{id:[0-9]+}", articleHandler)
	http.ListenAndServe(":9000", r)
}
