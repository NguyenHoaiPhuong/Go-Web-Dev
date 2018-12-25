package jsonfunc

import (
	"GO-WEB-DEV/101_Best-Practices/02_Practice/error"
	"encoding/json"
	"io"
	"os"
	"strings"
)

// ConvertFromJSON converts json object stored in reader and save into user object
func ConvertFromJSON(r io.Reader, object interface{}) error.Error {
	var err error.Imp
	osErr := json.NewDecoder(r).Decode(&object)
	if osErr != nil {
		err.InsertErrorMessage(osErr.Error())
		err.InsertErrorMessage(error.ErrorJSONConvert)
		return err
	}
	return nil
}

// ConvertToJSON converts object to JSON byte slice
func ConvertToJSON(object interface{}) ([]byte, error.Error) {
	var err error.Imp
	bs, osErr := json.Marshal(object)
	if osErr != nil {
		err.InsertErrorMessage(osErr.Error())
		err.InsertErrorMessage(error.ErrorJSONConvert)
		return nil, err
	}
	return bs, nil
}

// WriteToFile reads content stored in io.Reader, then write the content into json file with given fileName
func WriteToFile(r io.Reader, fileName string) error.Error {
	var err error.Imp
	ss := strings.Split(fileName, ".")
	if ss[len(ss)-1] != "json" {
		err.SetErrorMessage(error.ErrorFileExtension)
		err.InsertErrorMessage(error.ErrorJSONWrite)
		return err
	}

	file, osErr := os.Open(fileName)
	if osErr != nil {
		err.SetErrorMessage(osErr.Error())
		err.InsertErrorMessage(error.ErrorJSONWrite)
		return err
	}
	defer file.Close()

	_, osErr = io.Copy(file, r)
	if osErr != nil {
		err.SetErrorMessage(osErr.Error())
		err.InsertErrorMessage(error.ErrorJSONWrite)
		return err
	}

	return nil
}

// ReadFromFile reads json file with given fileName, then save into object
func ReadFromFile(fileName string, v interface{}) error.Error {
	var err error.Imp
	ss := strings.Split(fileName, ".")
	if ss[len(ss)-1] != "json" {
		err.SetErrorMessage(error.ErrorFileExtension)
		err.InsertErrorMessage(error.ErrorJSONRead)
		return err
	}

	file, osErr := os.Open(fileName)
	if osErr != nil {
		err.SetErrorMessage(osErr.Error())
		err.InsertErrorMessage(error.ErrorJSONRead)
		return err
	}
	defer file.Close()

	osErr = json.NewDecoder(file).Decode(&v)
	if osErr != nil {
		err.SetErrorMessage(osErr.Error())
		err.InsertErrorMessage(error.ErrorJSONRead)
		return err
	}
	return nil
}
