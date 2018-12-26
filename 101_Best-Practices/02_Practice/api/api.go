package api

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// API struct includes Router and Handler functions
type API struct {
	Router *mux.Router
}

// InitRouter initializes the router
func (api *API) InitRouter() {
	api.Router = mux.NewRouter()
}

// RegisterHandleFunction register the handle function to the router
func (api *API) RegisterHandleFunction(method string, path string, f func(w http.ResponseWriter, r *http.Request)) {
	method = strings.ToUpper(method)
	switch method {
	case "GET":
		api.get(path, f)
	case "PUT":
		api.put(path, f)
	case "POST":
		api.post(path, f)
	case "DELETE":
		api.delete(path, f)
	}
}

// get wraps the router for GET method
func (api *API) get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	api.Router.HandleFunc(path, f).Methods("GET")
}

// post wraps the router for POST method
func (api *API) post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	api.Router.HandleFunc(path, f).Methods("POST")
}

// put wraps the router for PUT method
func (api *API) put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	api.Router.HandleFunc(path, f).Methods("PUT")
}

// delete wraps the router for DELETE method
func (api *API) delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	api.Router.HandleFunc(path, f).Methods("DELETE")
}
