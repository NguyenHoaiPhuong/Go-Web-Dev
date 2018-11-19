package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json: "id"`
	Isbn   string  `json: "isbn"`
	Title  string  `json: "title"`
	Author *Author `json: "author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var books []Book

func init() {
	books = append(
		books,
		Book{ID: "1001", Isbn: "4865972", Title: "Book One", Author: &Author{"Nguyen", "Akagi"}},
		Book{ID: "1002", Isbn: "4951972", Title: "Book Two", Author: &Author{"Vo", "Yushin"}},
		Book{ID: "1003", Isbn: "3249972", Title: "Book Three", Author: &Author{"Nguyen", "Tien"}},
	)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Home Page</h1>")
}

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for _, book := range books {
		if book.ID == param["id"] {
			json.NewEncoder(w).Encode(&book)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	id := rand.Intn(100000000)
	book.ID = strconv.Itoa(id)
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func updateBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for idx, book := range books {
		if book.ID == param["id"] {
			var b Book
			json.NewDecoder(r.Body).Decode(&b)
			books[idx].Isbn = b.Isbn
			books[idx].Title = b.Title
			books[idx].Author = b.Author
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for idx, book := range books {
		if book.ID == param["id"] {
			books = append(books[:idx], books[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/books", getBooksHandler).Methods("GET")
	r.HandleFunc("/books/{id:[0-9]+}", getBookHandler).Methods("GET")
	r.HandleFunc("/books", createBookHandler).Methods("POST")
	r.HandleFunc("/books/{id:[0-9]+}", updateBookHandler).Methods("PUT")
	r.HandleFunc("/books/{id:[0-9]+}", deleteBookHandler).Methods("DELETE")

	srv := &http.Server{
		Addr:         "127.0.0.1:9000",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
