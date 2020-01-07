package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/016_Machinery/03_Example/server"
	"github.com/gin-gonic/gin"
)

var mode string

func main() {
	flag.StringVar(&mode, "mode", "master", "Default mode is master. Another available mode is worker")
	flag.Parse()
	log.Println("Running Mode:", mode)

	server.InitServer("./config/config.yml")

	switch mode {
	case "master":
		router := gin.Default()
		router.RedirectTrailingSlash = true
		router.RedirectFixedPath = true
		router.POST("/tasks", CreateTask)
		srv := &http.Server{
			Addr:    ":9000",
			Handler: router,
		}
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	case "worker":
		server.LaunchWorkerPool("yushin", 10)
		server.RegisterTasks()
		// Nothing in the startwroker pool is blocking. Need to wait to avoid ending the process.
		for {
			time.Sleep(1000 * time.Millisecond)
		}
	default:
		log.Fatalln("Unknown running mode")
	}
}
