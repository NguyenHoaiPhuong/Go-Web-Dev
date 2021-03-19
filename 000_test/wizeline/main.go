package main

import (
	"fmt"
	"time"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/000_test/wizeline/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	host = "localhost"
	port = "8000"
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
	if err := router.Run(fmt.Sprintf("%s:%s", host, port)); err != nil {
		fmt.Printf("router.Run error: %s\n", err.Error())
	}
}
