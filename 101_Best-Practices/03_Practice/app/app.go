package app

import (
	"log"
	"net/http"
	"time"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/101_Best-Practices/03_Practice/api"
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/101_Best-Practices/03_Practice/config"
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/101_Best-Practices/03_Practice/handlers"
)

// NewApp returns new app interface
func NewApp() IApp {
	return &ImpApp{}
}

// IApp : app interface
type IApp interface {
	Init()
	GetConfig() *config.Config
	Run()
}

// ImpApp struct includes router and mongodb session
type ImpApp struct {
	conf *config.Config
	srv  *api.Server

	IApp
}

// Init : initializes configurations, database, etc
func (app *ImpApp) Init() {
	app.initConfig()
	app.initServer()
}

func (app *ImpApp) initConfig() {
	app.conf = config.ParseConfig()
}

func (app *ImpApp) initServer() {
	app.srv = new(api.Server)
	app.srv.InitRouter()
	app.srv.RegisterHandleFunction("POST", api.PathCreateUser, handlers.CreateAccount)
	app.srv.RegisterHandleFunction("POST", api.PathLogin, handlers.Authenticate)
	app.srv.RegisterHandleFunction("POST", api.PathCreateContact, handlers.CreateContact)
	app.srv.RegisterHandleFunction("GET", api.PathListContacts, handlers.GetContactsFor)
}

// GetConfig : returns config
func (app *ImpApp) GetConfig() *config.Config {
	return app.conf
}

// Run App
func (app *ImpApp) Run() {
	srv := &http.Server{
		Handler:      app.srv.GetRouter(),
		Addr:         app.conf.Host + ":" + app.conf.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
