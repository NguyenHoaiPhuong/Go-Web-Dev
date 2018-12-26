package main

import "Go-Web-Dev/101_Best-Practices/02_Practice/app"

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run()
}
