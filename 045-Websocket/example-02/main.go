package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	go h.run()

	router := gin.New()
	router.LoadHTMLFiles("index.html")

	router.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomID := c.Param("roomId")
		serveWs(c.Writer, c.Request, roomID)
	})

	router.Run("0.0.0.0:8000")
}
