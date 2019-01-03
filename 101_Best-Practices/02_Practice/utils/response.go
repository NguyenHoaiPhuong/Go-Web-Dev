package utils

import (
	"Go-Web-Dev/101_Best-Practices/02_Practice/error"
	"encoding/json"
	"net/http"
)

// RespondError writes the error messages
func RespondError(w http.ResponseWriter, status int, messages ...string) {
	errNew := error.Imp{}
	for _, msg := range messages {
		errNew.InsertErrorMessage(msg)
	}
	w.WriteHeader(status)
	w.Write([]byte(errNew.Error()))
}

// RespondJSON writes the object
func RespondJSON(w http.ResponseWriter, status int, object interface{}) error.Error {
	bs, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		var errNew error.Imp
		errNew.SetErrorMessage(err.Error())
		return errNew
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(bs)
	return nil
}
