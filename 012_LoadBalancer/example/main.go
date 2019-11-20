package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	u, _ := url.Parse("http://localhost:8080")
	rp := httputil.NewSingleHostReverseProxy(u)
	// initialize your server and add this as handler
	http.HandlerFunc(rp.ServeHTTP)
}
