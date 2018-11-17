package main

import (
	"fmt"
	"net/http"
)

type handle_func struct{}

func (h handle_func) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the handler function")

	fmt.Fprintf(w, "Response Header:\n%s\n", w.Header())

	fmt.Fprintf(w, "Request Header:\n%v\n", r.Header)
	fmt.Fprintf(w, "Request Method:\n%s\n", r.Method)
	fmt.Fprintf(w, "Request Host:\n%s\n", r.Host)
	fmt.Fprintf(w, "Request Form:\n%v\n", r.Form)
	fmt.Fprintf(w, "Request Post Form:\n%v\n", r.PostForm)
	fmt.Fprintf(w, "Request Context:\n%v\n", r.Context())
}

func main() {
	var h handle_func
	http.ListenAndServe(":9000", h)
}
