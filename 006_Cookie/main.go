package main

import (
	"log"
	"net/http"
	"time"
)

// Handler : handler for homepage
type Handler struct {
}

func (m Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		log.Println("Write Cookie")
		expiration := time.Now().Add(2 * time.Minute) // expire in 2 minutes
		cookie := http.Cookie{Name: "username", Value: "akagi", Expires: expiration}
		http.SetCookie(w, &cookie)
	case "/cookie":
		log.Println("Read Cookie")
		cookie, err := r.Cookie("username")
		if err != nil {
			log.Println("Get Cookie Error:", err.Error())
		}
		log.Println("Cookie:", cookie)
		log.Println("Name:", cookie.Name)
		log.Println("Value:", cookie.Value)
	}
}

func main() {
	var h Handler
	log.Fatalln(http.ListenAndServe(":9000", h))
}
