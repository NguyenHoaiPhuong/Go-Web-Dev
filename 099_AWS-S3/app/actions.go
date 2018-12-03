package app

type appActions interface {
	Init()
	Run()
}

type s3Actions interface {
	AWSS3ListBuckets()
	AWSS3ListBucketItems(string)
	AWSS3CreateNewBucket(string)
	AWSS3UploadFileToBucket(string, string)
}
