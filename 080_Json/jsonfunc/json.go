package jsonfunc

import (
	"encoding/json"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
)

// WriteToFile reads content stored in io.Reader, then write the content into json file with given fileName
func WriteToFile(r io.Reader, fileName string) error {
	var err error
	ss := strings.Split(fileName, ".")
	if ss[len(ss)-1] != "json" {
		err = errors.New("File Extension Mismatched")
		return err
	}

	file, osErr := os.Create(fileName)
	if osErr != nil {
		return osErr
	}
	defer file.Close()
	io.Copy(file, r)
	return nil
}

// ReadFromFile reads json file with given fileName, then save into object
func ReadFromFile(fileName string, v interface{}) error {
	var err error
	ss := strings.Split(fileName, ".")
	if ss[len(ss)-1] != "json" {
		err = errors.New("File Extension Mismatched")
		return err
	}

	file, osErr := os.Open(fileName)
	if osErr != nil {
		return osErr
	}
	defer file.Close()
	osErr = json.NewDecoder(file).Decode(&v)
	if osErr != nil {
		return osErr
	}
	return nil
}

// ConvertToJSON converts object to JSON byte slice
func ConvertToJSON(object interface{}) ([]byte, error) {
	bs, osErr := json.Marshal(object)
	if osErr != nil {
		return nil, osErr
	}
	return bs, nil
}

// ConvertFromJSON converts json object stored in reader and save into user object
func ConvertFromJSON(r io.Reader, object interface{}) error {
	osErr := json.NewDecoder(r).Decode(&object)
	if osErr != nil {
		return osErr
	}
	return nil
}
