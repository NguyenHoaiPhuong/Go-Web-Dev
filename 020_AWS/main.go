package main

import (
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/020_AWS/config"
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/020_AWS/services"
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/020_AWS/services/awss3"
)

func main() {
	conf := config.GlobalConfig()
	conf.Print()

	svc := new(services.Service)
	svc.Init()

	TestS3(svc.S3)
	TestEC2()
}

// TestS3 test functions in S3 service
func TestS3(svc awss3.IService) {
	// bucket1 := "akagi1"
	bucket2 := "akagi2"
	// svc.CreateNewBucket(bucket1)
	svc.CreateNewBucket(bucket2)
	// svc.ListBuckets()
	// file := "resources/config.json"
	// svc.UploadFileToBucket(file, bucket1)
	// svc.ListBucketItems(bucket1)
	// svc.CopyItemFromBucketToBucket(bucket1, bucket2, file)
	svc.UploadDirectoryToBucket("./resources", bucket2, "resources")
	svc.ListBucketItems(bucket2)
	// svc.DownloadFileFromBucket("resource/CV_Nguyen-Hoai-Phuong_181203.docx", "akagi21061986")
	// svc.DeleteAllBucketItems("akagi21061986")
	svc.DeleteBucket("akagi21061986")
}

// TestEC2 test functions in EC2 service
func TestEC2() {

}
