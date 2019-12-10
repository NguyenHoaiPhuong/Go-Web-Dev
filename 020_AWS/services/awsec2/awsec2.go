package awsecs

import "github.com/NguyenHoaiPhuong/Go-Web-Dev/020_AWS/config"

// IService includes all methods related to S3
type IService interface {
	init(cf config.S3Configurations)

	CreateNewBucket(string) error
	DeleteBucket(string) error
	ListBuckets() error
	ListBucketItems(string) error
	DeleteBucketItem(string, string) error
	DeleteAllBucketItems(string) error
	UploadFileToBucket(string, string) error
	DownloadFileFromBucket(string, string) error
	CopyItemFromBucketToBucket(string, string, string) error
}
