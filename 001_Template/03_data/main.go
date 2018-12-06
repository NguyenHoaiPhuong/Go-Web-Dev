package main

import (
	"log"
	"os"
	"text/template"
)

// Account struct
type Account struct {
	Name string
	Pass string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("html/*.html"))
}
func main() {
	acc := Account{
		Name: "akagi",
		Pass: "1234",
	}

	// Single variable
	err := tpl.ExecuteTemplate(os.Stdout, "data.html", 99)
	if err != nil {
		log.Fatalln(err)
	}

	// Struct
	err = tpl.ExecuteTemplate(os.Stdout, "struct.html", acc)
	if err != nil {
		log.Fatalln(err)
	}
}
