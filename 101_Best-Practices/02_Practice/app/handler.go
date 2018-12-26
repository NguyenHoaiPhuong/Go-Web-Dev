package app

import (
	"Go-Web-Dev/101_Best-Practices/02_Practice/error"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func respondError(w http.ResponseWriter, status int, messages ...string) {
	errNew := error.Imp{}
	for _, msg := range messages {
		errNew.InsertErrorMessage(msg)
	}
	w.WriteHeader(status)
	w.Write([]byte(errNew.Error()))
}

func respondJSON(w http.ResponseWriter, status int, object interface{}) error.Error {
	bs, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		var errNew error.Imp
		errNew.SetErrorMessage(err.Error())
		return errNew
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(bs)
	return nil
}

func (a *App) allBooks(w http.ResponseWriter, r *http.Request) {
	books, err := a.Database.GetAllDocuments(*a.Config.MongoDBConfig.Name, *a.Config.MongoDBConfig.Collection)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Database error")
		var errNew error.Imp
		errNew.SetErrorMessage(err.Error())
		errNew.InsertErrorMessage(error.ErrorAppGetAllBooks)
		log.Printf("%v\n", errNew.Error())
		return
	}

	err = respondJSON(w, http.StatusOK, books)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "JSON error")
		var errNew error.Imp
		errNew.SetErrorMessage(err.Error())
		errNew.InsertErrorMessage(error.ErrorAppGetAllBooks)
		log.Printf("%v\n", errNew.Error())
		return
	}
}

func (a *App) addBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "addBook\n")
}

func (a *App) bookByISBN(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isbn := vars["isbn"]
	book, err := a.Database.GetDocumentByKey(*a.Config.MongoDBConfig.Name, *a.Config.MongoDBConfig.Collection, "isbn", isbn)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "Database error")
		var errNew error.Imp
		errNew.SetErrorMessage(err.Error())
		errNew.InsertErrorMessage(error.ErrorAppGetBookByIsbn)
		log.Printf("%v\n", errNew.Error())
		return
	}

	err = respondJSON(w, http.StatusOK, book)
	if err != nil {
		respondError(w, http.StatusInternalServerError, "JSON error")
		var errNew error.Imp
		errNew.SetErrorMessage(err.Error())
		errNew.InsertErrorMessage(error.ErrorAppGetBookByIsbn)
		log.Printf("%v\n", errNew.Error())
		return
	}
}

func (a *App) updateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "updateBook\n")
}

func (a *App) deleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "deleteBook\n")
}
