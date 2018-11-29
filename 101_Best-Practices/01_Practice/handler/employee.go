package handler

import (
	"GO-WEB-DEV/101_Best-Practices/01_Practice/jsonfunc"
	"GO-WEB-DEV/101_Best-Practices/01_Practice/model"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
)

// GetAllEmployees gets all employees' info in the database
func GetAllEmployees(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	employees := []model.Employee{}
	db.Find(&employees)
	respondJSON(w, employees)
}

// GetEmployee gets the specific employee s' info in the database
func GetEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["ID"]
	employee := new(model.Employee)
	db.First(employee, id)
	if employee == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	respondJSON(w, employee)
}

// CreateEmployee creates new employee
func CreateEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

// UpdateEmployee update specific employee s' info in the database
func UpdateEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

// DeleteEmployee deletes specific employee s' info in the database
func DeleteEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

// EnableEmployee enables specific employee in the database
func EnableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

// DisableEmployee disables specific employee in the database
func DisableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}

func respondJSON(w http.ResponseWriter, object interface{}) {
	bs, err := jsonfunc.ConvertToJSON(object)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}
