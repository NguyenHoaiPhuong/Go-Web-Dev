package api

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Server struct includes Router and Handler functions
type Server struct {
	router *mux.Router
}

// InitRouter initializes the router
func (srv *Server) InitRouter() {
	srv.router = mux.NewRouter()
	srv.router.Use(JwtAuthentication) //attach JWT auth middleware
}

// GetRouter returns the server router
func (srv *Server) GetRouter() *mux.Router {
	return srv.router
}

// RegisterHandleFunction register the handle function to the router
func (srv *Server) RegisterHandleFunction(method string, path string, f func(w http.ResponseWriter, r *http.Request)) {
	method = strings.ToUpper(method)
	switch method {
	case "GET":
		srv.get(path, f)
	case "PUT":
		srv.put(path, f)
	case "POST":
		srv.post(path, f)
	case "DELETE":
		srv.delete(path, f)
	}
}

// get wraps the router for GET method
func (srv *Server) get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	srv.router.HandleFunc(path, f).Methods("GET")
}

// post wraps the router for POST method
func (srv *Server) post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	srv.router.HandleFunc(path, f).Methods("POST")
}

// put wraps the router for PUT method
func (srv *Server) put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	srv.router.HandleFunc(path, f).Methods("PUT")
}

// delete wraps the router for DELETE method
func (srv *Server) delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	srv.router.HandleFunc(path, f).Methods("DELETE")
}
