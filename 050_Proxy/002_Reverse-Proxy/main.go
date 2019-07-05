package main

import (
	"net/http"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/050_Proxy/002_Reverse-Proxy/proxy"
)

func main() {
	addr := ":8000"
	prox := new(proxy.ReverseProxy)
	http.Handle("/", prox)
	http.ListenAndServe(addr, prox)
}
