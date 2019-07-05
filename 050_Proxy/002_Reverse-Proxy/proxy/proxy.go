package proxy

import (
	"fmt"
	"log"
	"net/http"
)

// ReverseProxy : load balancing
type ReverseProxy struct {
}

func (proxy *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// http://localhost:9090: prometheus
	// http://localhost:9091: betta-house
	log.Println("Method:", r.Method)
	log.Println("Remote Address:", r.RemoteAddr)
	log.Println("Host:", r.Host)
	log.Println("RequestURI:", r.RequestURI)

	log.Println("URL Host:", r.URL.Host)
	log.Println("URL Hostname:", r.URL.Hostname())
	log.Println("URL Path:", r.URL.Path)

	fmt.Println("-----------------------------")

	if r.RequestURI == "/prometheus" {
		http.Get("http://localhost:9090")
	} else if r.RequestURI == "/betta" {
		http.Get("http://localhost:9091")
	}
}
