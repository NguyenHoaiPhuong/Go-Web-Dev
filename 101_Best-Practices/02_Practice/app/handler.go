package app

import (
	"Go-Web-Dev/101_Best-Practices/02_Practice/error"
	"Go-Web-Dev/101_Best-Practices/02_Practice/jsonfunc"
	"Go-Web-Dev/101_Best-Practices/02_Practice/model"
	"Go-Web-Dev/101_Best-Practices/02_Practice/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *App) allBooks(w http.ResponseWriter, r *http.Request) {
	books, err := a.Database.GetAllDocuments(*a.Config.MongoDBConfig.Name, *a.Config.MongoDBConfig.Collection)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, error.ErrorDB)
		var errNew error.Imp
		errNew.SetErrorMessage(err.Error())
		errNew.InsertErrorMessage(error.ErrorAppGetAllBooks)
		log.Printf("%v\n", errNew.Error())
		return
	}

	err = utils.RespondJSON(w, http.StatusOK, books)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, error.ErrorJSON)
		var errNew error.Imp
		errNew.SetErrorMessage(err.Error())
		errNew.InsertErrorMessage(error.ErrorAppGetAllBooks)
		log.Printf("%v\n", errNew.Error())
		return
	}
}

func (a *App) addBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book

	err := jsonfunc.ConvertFromJSON(r.Body, &book)
	defer r.Body.Close()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, error.ErrorJSON)
		var errNew error.Imp
		errNew.SetErrorMessage(err.Error())
		errNew.InsertErrorMessage(error.ErrorAppAddBook)
		log.Printf("%v\n", errNew.Error())
		return
	}

	err = a.Database.AddDocument(*a.Config.MongoDBConfig.Name, *a.Config.MongoDBConfig.Collection, &book)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, error.ErrorDB)
		var errNew error.Imp
		errNew.SetErrorMessage(err.Error())
		errNew.InsertErrorMessage(error.ErrorAppAddBook)
		log.Printf("%v\n", errNew.Error())
		return
	}

	fmt.Fprintf(w, "Add new book into database successfully\n")
	utils.RespondJSON(w, http.StatusOK, book)
}

func (a *App) bookByISBN(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isbn := vars["isbn"]
	book, err := a.Database.GetDocumentByKey(*a.Config.MongoDBConfig.Name, *a.Config.MongoDBConfig.Collection, "isbn", isbn)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, error.ErrorDB)
		var errNew error.Imp
		errNew.SetErrorMessage(err.Error())
		errNew.InsertErrorMessage(error.ErrorAppGetBookByIsbn)
		log.Printf("%v\n", errNew.Error())
		return
	}

	err = utils.RespondJSON(w, http.StatusOK, book)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, error.ErrorJSON)
		var errNew error.Imp
		errNew.SetErrorMessage(err.Error())
		errNew.InsertErrorMessage(error.ErrorAppGetBookByIsbn)
		log.Printf("%v\n", errNew.Error())
		return
	}
}

func (a *App) updateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isbn := vars["isbn"]
	var book model.Book
	err := jsonfunc.ConvertFromJSON(r.Body, &book)
	defer r.Body.Close()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, error.ErrorJSON)
		var errNew error.Imp
		errNew.SetErrorMessage(err.Error())
		errNew.InsertErrorMessage(error.ErrorAppUpdateBook)
		log.Printf("%v\n", errNew.Error())
		return
	}
	book.ISBN = isbn
	err = a.Database.UpdateDocument(*a.Config.MongoDBConfig.Name, *a.Config.MongoDBConfig.Collection, &book)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, error.ErrorDB)
		var errNew error.Imp
		errNew.SetErrorMessage(err.Error())
		errNew.InsertErrorMessage(error.ErrorAppUpdateBook)
		log.Printf("%v\n", errNew.Error())
		return
	}
	fmt.Fprintf(w, "Update book with isbn %v into database successfully\n", isbn)
	utils.RespondJSON(w, http.StatusOK, book)
}

func (a *App) deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	isbn := vars["isbn"]
	err := a.Database.DeleteDocumentByKey(*a.Config.MongoDBConfig.Name, *a.Config.MongoDBConfig.Collection, "isbn", isbn)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, error.ErrorDB)
		var errNew error.Imp
		errNew.SetErrorMessage(err.Error())
		errNew.InsertErrorMessage(error.ErrorAppDeleteBookByIsbn)
		log.Printf("%v\n", errNew.Error())
		return
	}

	fmt.Fprintf(w, "Remove book with isbn %v from database successfully\n", isbn)
}
