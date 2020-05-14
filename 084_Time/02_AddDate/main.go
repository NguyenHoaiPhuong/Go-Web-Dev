package main

import (
	"log"
	"time"
)

func main() {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, "2020-01-30 23:59:59")
	if err != nil {
		log.Fatal(err)
	}
	t = t.AddDate(0, 1, 0)
	// Jan 03th + 1 month = Mar 01st
	log.Println(t)
}
