package main

import (
	"GO-WEB-DEV/101_Best-Practices/01_Practice/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run()
}
