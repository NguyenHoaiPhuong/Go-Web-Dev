package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Server : struct
type Server struct {
	router *gin.Engine
}

// NewServer : returns a ptr of server struct
func NewServer(router *gin.Engine) *Server {
	return &Server{
		router: router,
	}
}

// Routes : sets groups of routes
func (s *Server) Routes() {
	s.router.GET("/", Homepage)
	api := s.router.Group("/api")
	{
		api.GET("/", RestAPIHomepage)
	}
	apiUser := api.Group("/user")
	{
		apiUser.POST("/register", CreateUser)
	}
}

// Run the server
func (s *Server) Run(address string) {
	if err := s.router.Run(address); err != nil {
		log.Fatalln(err)
	}
}
