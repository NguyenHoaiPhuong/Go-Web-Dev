package main

import (
	"fmt"
	"net/http"
)

func hello(db Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, db.URL)
	}
}
