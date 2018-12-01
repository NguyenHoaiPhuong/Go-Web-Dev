package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var tpl = &template.Template{}

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	http.HandleFunc("/", uploadFile)
	http.ListenAndServe(":9000", nil)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)

	var s string
	if r.Method == http.MethodPost {
		f, _, err := r.FormFile("usrfile")
		if err != nil {
			log.Println(err)
			http.Error(w, "Error upload file", http.StatusInternalServerError)
			return
		}
		defer f.Close()

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error read file", http.StatusInternalServerError)
			return
		}

		s = string(bs)
	}
	fmt.Fprint(w, s)
}
