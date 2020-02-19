package main

import (
	"fmt"
	"net/http"
	"time"
)

func timed(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		time.Sleep(time.Millisecond)
		f(w, r)
		end := time.Now()
		fmt.Println("The request took", end.Sub(start).Nanoseconds())
	}
}

func main() {
	http.HandleFunc("/hello", timed(hello))
	http.ListenAndServe(":9000", nil)
}
