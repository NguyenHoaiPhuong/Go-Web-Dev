package main

import "net/http"

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../client/index.html")
}
