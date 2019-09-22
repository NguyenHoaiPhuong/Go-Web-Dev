package awss3

import (
	"github.com/NguyenHoaiPhuong/GO-WEB-DEV/099_AWS-S3/config"
)

// Actions includes all functions related to S3
type Actions interface {
	Init(cf *config.S3Config)

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
