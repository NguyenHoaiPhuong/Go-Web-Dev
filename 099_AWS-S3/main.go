package main

import (
	"GO-WEB-DEV/099_AWS-S3/app"
)

func main() {
	app := new(app.App)
	app.Init()

	app.S3Service.ListBuckets()
	app.S3Service.CreateNewBucket("akagi")

	//app.Run()
}
