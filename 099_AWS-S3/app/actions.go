package app

type appActions interface {
	Init()
	Run()
}

type s3Actions interface {
	AWSS3CreateNewBucket(string)
	AWSS3DeleteBucket(string)
	AWSS3ListBuckets()
	AWSS3ListBucketItems(string)
	AWSS3DeleteBucketItem(string)
	AWSS3DeleteAllBucketItems(string)
	AWSS3UploadFileToBucket(string, string)
	AWSS3DownloadFileFromBucket(string, string)
}
