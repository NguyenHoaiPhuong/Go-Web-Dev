package app

import (
	"Go-Web-Dev/101_Best-Practices/02_Practice/api"
	"Go-Web-Dev/101_Best-Practices/02_Practice/config"
	"Go-Web-Dev/101_Best-Practices/02_Practice/error"
	"Go-Web-Dev/101_Best-Practices/02_Practice/repo"
	"log"
	"net/http"
	"time"
)

// App struct includes router and mongodb session
type App struct {
	Config   *config.Config
	Database *repo.MongoDB
	API      *api.API
}

// Run App
func (a *App) Run() {
	srv := &http.Server{
		Handler:      a.API.Router,
		Addr:         ":9000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

	defer a.Database.Session.Close()
}

// Initialize init App
func (a *App) Initialize() {
	a.initConfigure()
	a.initDatabase()
	a.initAPI()
}

func (a *App) initConfigure() {
	a.Config = config.SetupConfig("resource/config.json")
}

func (a *App) initDatabase() {
	a.Database = new(repo.MongoDB)
	err := a.Database.InitDBSession(*a.Config.MongoDBConfig.Host)
	if err != nil {
		var errNew error.Imp
		errNew.InsertErrorMessage(err.Error())
		errNew.InsertErrorMessage(error.ErrorAppInit)
		log.Fatalln(errNew.Error())
	}
	a.Database.EnsureIndex("store", "books", "isbn")
}

func (a *App) initAPI() {
	a.API = new(api.API)
	a.API.InitRouter()
	a.API.RegisterHandleFunction("GET", "/books", a.allBooks)
	a.API.RegisterHandleFunction("GET", "/book/{isbn}", a.bookByISBN)
	a.API.RegisterHandleFunction("POST", "/books", a.addBook)
	a.API.RegisterHandleFunction("PUT", "/book/{isbn}", a.updateBook)
	a.API.RegisterHandleFunction("DELETE", "/book/{isbn}", a.deleteBook)
}
