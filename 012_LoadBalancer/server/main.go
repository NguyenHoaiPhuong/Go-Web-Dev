package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

type handleFunc struct{}

func (h handleFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the handler function")
	fmt.Fprintln(w, "Host:", host)
	fmt.Fprintln(w, "Port:", port)

	fmt.Fprintf(w, "Response Header:\n%s\n", w.Header())

	fmt.Fprintf(w, "Request Header:\n%v\n", r.Header)
	fmt.Fprintf(w, "Request Method:\n%s\n", r.Method)
	fmt.Fprintf(w, "Request Host:\n%s\n", r.Host)
	fmt.Fprintf(w, "Request Form:\n%v\n", r.Form)
	fmt.Fprintf(w, "Request Post Form:\n%v\n", r.PostForm)
	fmt.Fprintf(w, "Request Context:\n%v\n", r.Context())
}

var (
	host string
	port string
)

func init() {
	flag.StringVar(&host, "host", "localhost", "Server host")
	flag.StringVar(&port, "port", ":3001", "Server port")
	flag.Parse()
	log.Println("Server host:", host)
	log.Println("Server port:", port)
}

func main() {
	var h handleFunc
	log.Fatal(http.ListenAndServe(host+":"+port, h))
}
