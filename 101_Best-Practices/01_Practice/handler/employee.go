package handler

import (
	"GO-WEB-DEV/101_Best-Practices/01_Practice/error"
	"GO-WEB-DEV/101_Best-Practices/01_Practice/jsonfunc"
	"GO-WEB-DEV/101_Best-Practices/01_Practice/model"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
)

// GetAllEmployees gets all employees' info in the database
func GetAllEmployees(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	employees := []model.Employee{}
	osErr := db.Find(&employees).Error
	if osErr != nil {
		respondError(w, http.StatusInternalServerError, osErr.Error(), error.ErrorFindData)
	}
	respondJSON(w, http.StatusOK, employees)
}

// GetEmployee gets the specific employee s' info in the database
func GetEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	employee := new(model.Employee)
	err := db.First(employee, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		errNew := error.ErrorImp{}
		errNew.InsertErrorMessage(err.Error())
		errNew.InsertErrorMessage(error.ErrorFindData)
		w.Write([]byte(errNew.Error()))
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

// CreateEmployee creates new employee
func CreateEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	employee := new(model.Employee)
	err := jsonfunc.ConvertFromJSON(r.Body, employee)
	defer r.Body.Close()
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error(), error.ErrorCreateData)
		return
	}
	osErr := db.Create(employee).Error
	if osErr != nil {
		respondError(w, http.StatusInternalServerError, osErr.Error(), error.ErrorCreateData)
	}
	respondJSON(w, http.StatusOK, employee)
}

// UpdateEmployee update specific employee s' info in the database
func UpdateEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	employee := new(model.Employee)
	osErr := db.First(employee, id).Error
	if osErr != nil {
		respondError(w, http.StatusNotFound, osErr.Error(), error.ErrorUpdateData)
		return
	}
	err := jsonfunc.ConvertFromJSON(r.Body, employee)
	defer r.Body.Close()
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error(), error.ErrorUpdateData)
		return
	}
	employee.ID = id
	osErr = db.Save(employee).Error
	if osErr != nil {
		respondError(w, http.StatusInternalServerError, osErr.Error(), error.ErrorUpdateData)
		return
	}
}

// DeleteEmployee deletes specific employee s' info in the database
func DeleteEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	employee := new(model.Employee)
	osErr := db.First(employee, id).Error
	if osErr != nil {
		respondError(w, http.StatusNotFound, osErr.Error(), error.ErrorDeleteData)
		return
	}
	osErr = db.Delete(employee).Error
	if osErr != nil {
		respondError(w, http.StatusInternalServerError, osErr.Error(), error.ErrorDeleteData)
		return
	}
}

// EnableEmployee enables specific employee in the database
func EnableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	employee := new(model.Employee)
	osErr := db.First(employee, model.Employee{ID: id}).Error
	if osErr != nil {
		respondError(w, http.StatusNotFound, osErr.Error(), error.ErrorUpdateData)
		return
	}
	employee.Enable()
	osErr = db.Save(employee).Error
	if osErr != nil {
		respondError(w, http.StatusInternalServerError, osErr.Error(), error.ErrorUpdateData)
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

// DisableEmployee disables specific employee in the database
func DisableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	employee := new(model.Employee)
	osErr := db.First(employee, model.Employee{ID: id}).Error
	if osErr != nil {
		respondError(w, http.StatusNotFound, osErr.Error(), error.ErrorUpdateData)
		return
	}
	employee.Disable()
	osErr = db.Save(employee).Error
	if osErr != nil {
		respondError(w, http.StatusInternalServerError, osErr.Error(), error.ErrorUpdateData)
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

func respondJSON(w http.ResponseWriter, status int, object interface{}) {
	bs, err := jsonfunc.ConvertToJSON(object)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(bs)
}

func respondError(w http.ResponseWriter, status int, messages ...string) {
	errNew := error.ErrorImp{}
	for _, msg := range messages {
		errNew.InsertErrorMessage(msg)
	}
	w.WriteHeader(status)
	w.Write([]byte(errNew.Error()))
}
