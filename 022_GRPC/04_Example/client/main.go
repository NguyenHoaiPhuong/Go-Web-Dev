package main

import (
	"time"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/022_GRPC/04_Example/client/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	srvAddr = "localhost:9001"
	cliAddr = "localhost:9000"
)

func main() {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://*", "https://*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		MaxAge:           12 * time.Hour,
	}))

	svr := api.NewServer(router)
	svr.Routes()
	svr.Run(cliAddr)
}
