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
	appActions
	s3Actions

	config    *config.Config
	router    *mux.Router
	s3Service *awss3.Service
}

// Init initializes settings
func (a *App) Init() {
	a.initConfig()
	a.initRouter()
	a.initAWSS3()
}

func (a *App) initConfig() {
	var err error
	a.config, err = config.GetConfig("resource/config.json")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func (a *App) initRouter() {
	a.router = mux.NewRouter()
}

func (a *App) initAWSS3() {
	a.s3Service = new(awss3.Service)
	a.s3Service.Init(a.config.S3Config)
}

// AWSS3ListBuckets lists all buckets
func (a *App) AWSS3ListBuckets() {
	err := a.s3Service.ListBuckets()
	if err != nil {
		panic(err)
	}
}

// AWSS3ListBucketItems lists all objects in the buckets
func (a *App) AWSS3ListBucketItems(bucketName string) {
	err := a.s3Service.ListBucketItems(bucketName)
	if err != nil {
		panic(err)
	}
}

// AWSS3CreateNewBucket creates new bucket
func (a *App) AWSS3CreateNewBucket(bucketName string) {
	err := a.s3Service.CreateNewBucket(bucketName)
	if err != nil {
		panic(err)
	}
}

// AWSS3UploadFileToBucket uploads a file to bucket
func (a *App) AWSS3UploadFileToBucket(fileName string, bucketName string) {
	err := a.s3Service.UploadFileToBucket(fileName, bucketName)
	if err != nil {
		panic(err)
	}
}

// Run runs the web app server
func (a *App) Run() {
	srv := &http.Server{
		Handler:      a.router,
		Addr:         ":9000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
