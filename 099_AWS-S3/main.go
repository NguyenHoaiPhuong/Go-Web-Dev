package main

import (
	"github.com/NguyenHoaiPhuong/GO-WEB-DEV/099_AWS-S3/app"
)

func main() {
	app := new(app.App)
	app.Init()

	bucket1 := "akagi1"
	bucket2 := "akagi2"
	// app.AWSS3CreateNewBucket(bucket1)
	// app.AWSS3CreateNewBucket(bucket2)
	app.AWSS3ListBuckets()
	file := "resource/CV_Nguyen-Hoai-Phuong_181203.docx"
	// app.AWSS3UploadFileToBucket(file, bucket1)
	app.AWSS3ListBucketItems(bucket1)
	app.AWSS3CopyItemFromBucketToBucket(bucket1, bucket2, file)
	app.AWSS3ListBucketItems(bucket2)
	//app.AWSS3DownloadFileFromBucket("resource/CV_Nguyen-Hoai-Phuong_181203.docx", "akagi21061986")
	//app.AWSS3DeleteAllBucketItems("akagi21061986")
	//app.AWSS3DeleteBucket("akagi21061986")
	//app.Run()
}
