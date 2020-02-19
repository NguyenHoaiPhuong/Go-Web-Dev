package main

import (
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/101_Best-Practices/03_Practice/app"
)

func main() {
	app := app.NewApp()
	app.Init()
	app.Run()
}
