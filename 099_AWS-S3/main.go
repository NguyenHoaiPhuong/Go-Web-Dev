package main

import (
	"github.com/NguyenHoaiPhuong/GO-WEB-DEV/099_AWS-S3/app"
)

func main() {
	app := new(app.App)
	app.Init()

	app.AWSS3ListBuckets()
	app.AWSS3CreateNewBucket("akagi21061986")
	//app.AWSS3ListBucketItems("matchtalent.com")
	//app.AWSS3UploadFileToBucket("resource/CV_Nguyen-Hoai-Phuong_181203.docx", "akagi21061986")
	//app.AWSS3DownloadFileFromBucket("resource/CV_Nguyen-Hoai-Phuong_181203.docx", "akagi21061986")
	//app.AWSS3DeleteAllBucketItems("akagi21061986")
	//app.AWSS3DeleteBucket("akagi21061986")
	//app.Run()
}
