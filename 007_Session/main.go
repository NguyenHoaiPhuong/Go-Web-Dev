/*
Sessions use cookies to save session ids on the client side,
and save all other information on the server side
*/

package main

import (
	"log"
	"net/http"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/007_Session/session"
)

var globalSessions *session.Manager

// Then, initialize the session manager
func init() {
	session.Init()
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	// if err != nil {
	// 	log.Fatalln(fmt.Errorf("Parsing template file error : %s", err.Error()))
	// }
	go globalSessions.GC()
}

func main() {
	http.HandleFunc("/", sayhelloName) // setting router rule
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9000", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
