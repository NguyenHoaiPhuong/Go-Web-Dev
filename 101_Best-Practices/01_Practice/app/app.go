package app

import (
	"GO-WEB-DEV/101_Best-Practices/01_Practice/error"
	"GO-WEB-DEV/101_Best-Practices/01_Practice/handler"
	"GO-WEB-DEV/101_Best-Practices/01_Practice/repo"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
	Rooter *mux.Router
	DB     *gorm.DB
	Host   string
}

// Run App
func (a *App) Run() {
	srv := &http.Server{
		Handler:      a.Rooter,
		Addr:         ":9000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

// Initialize init App
func (a *App) Initialize() {
	a.initDatabase()
	a.initRouter()
}

func (a *App) initDatabase() {
	db, err := repo.GetDatabase()
	if err != nil {
		var errNew error.ErrorImp
		errNew.InsertErrorMessage(err.Error())
		errNew.InsertErrorMessage(error.ErrorApplicationInit)
		log.Fatalln(errNew.Error())
	}
	a.DB = db
}

func (a *App) initRouter() {
	a.Rooter = mux.NewRouter()
	a.Get("/employees", a.GetAllEmployees)
	a.Post("/employees", a.CreateEmployee)
	a.Get("/employees/{id}", a.GetEmployee)
	a.Put("/employees/{id}", a.UpdateEmployee)
	a.Delete("/employees/{id}", a.DeleteEmployee)
	a.Put("/employees/{id}/enable", a.EnableEmployee)
	a.Put("/employees/{id}/disable", a.DisableEmployee)
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Rooter.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Rooter.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Rooter.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Rooter.HandleFunc(path, f).Methods("DELETE")
}

// GetAllEmployees gets all employees' info in the database
func (a *App) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	handler.GetAllEmployees(a.DB, w, r)
}

// GetEmployee gets the specific employee s' info in the database
func (a *App) GetEmployee(w http.ResponseWriter, r *http.Request) {
	handler.GetEmployee(a.DB, w, r)
}

// CreateEmployee creates new employee
func (a *App) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	handler.CreateEmployee(a.DB, w, r)
}

// UpdateEmployee update specific employee s' info in the database
func (a *App) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	handler.UpdateEmployee(a.DB, w, r)
}

// DeleteEmployee deletes specific employee s' info in the database
func (a *App) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	handler.DeleteEmployee(a.DB, w, r)
}

// EnableEmployee enables specific employee in the database
func (a *App) EnableEmployee(w http.ResponseWriter, r *http.Request) {
	handler.EnableEmployee(a.DB, w, r)
}

// DisableEmployee disables specific employee in the database
func (a *App) DisableEmployee(w http.ResponseWriter, r *http.Request) {
	handler.DisableEmployee(a.DB, w, r)
}
