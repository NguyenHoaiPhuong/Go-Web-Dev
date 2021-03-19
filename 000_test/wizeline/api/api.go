package api

import "github.com/gin-gonic/gin"

// Server : struct
type Server struct {
	g *gin.Engine
}

// NewServer :
func NewServer(g *gin.Engine) *Server {
	return &Server{
		g: g,
	}
}
