package app

import (
	"GO-WEB-DEV/099_AWS-S3/awss3"
	"GO-WEB-DEV/099_AWS-S3/config"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// App implements actions
type App struct {
	Config    *config.Config
	Router    *mux.Router
	S3Service *awss3.Service
}

// Init initializes settings
func (a *App) Init() {
	a.initConfig()
	a.initRouter()
	a.initAWSS3()
}

func (a *App) initConfig() {
	var err error
	a.Config, err = config.GetConfig("resource/config.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func (a *App) initRouter() {
	a.Router = mux.NewRouter()
}

func (a *App) initAWSS3() {
	a.S3Service = new(awss3.Service)
	a.S3Service.Init(a.Config.S3Config)
}

// Run runs the web app server
func (a *App) Run() {
	srv := &http.Server{
		Handler:      a.Router,
		Addr:         ":9000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
