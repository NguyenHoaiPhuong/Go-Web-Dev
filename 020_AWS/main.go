package main

import (
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/020_AWS/config"
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/020_AWS/services"
)

func main() {
	conf := config.GlobalConfig()
	conf.Print()

	svc := new(services.Service)
	svc.Init()

	bucket1 := "akagi1"
	bucket2 := "akagi2"
	svc.S3.CreateNewBucket(bucket1)
	svc.S3.CreateNewBucket(bucket2)
	svc.S3.ListBuckets()
	file := "resources/config.json"
	svc.S3.UploadFileToBucket(file, bucket1)
	svc.S3.ListBucketItems(bucket1)
	svc.S3.CopyItemFromBucketToBucket(bucket1, bucket2, file)
	svc.S3.ListBucketItems(bucket2)
	// svc.S3.DownloadFileFromBucket("resource/CV_Nguyen-Hoai-Phuong_181203.docx", "akagi21061986")
	// svc.S3.DeleteAllBucketItems("akagi21061986")
	// svc.S3.DeleteBucket("akagi21061986")
}
