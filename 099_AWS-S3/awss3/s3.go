package awss3

import (
	"GO-WEB-DEV/099_AWS-S3/config"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Service implements all methods in the Actions interface
type Service struct {
	Actions

	Client   *s3.S3
	Uploader *s3manager.Uploader
}

// Init initializes settings
func (svc *Service) Init(s3config *config.S3Config) {
	config := &aws.Config{
		Region: aws.String(s3config.Region),
	}
	sess := session.Must(session.NewSession(config))
	svc.Client = s3.New(sess)
	svc.Uploader = s3manager.NewUploader(sess)
}

// CreateNewBucket creates new bucket
func (svc *Service) CreateNewBucket(bucketName string) error {
	_, err := svc.Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create bucket %s, %v\n", bucketName, err)
		return err
	}

	// Wait until bucket is created before finishing
	fmt.Printf("Waiting for bucket %q to be created...\n", bucketName)

	err = svc.Client.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred while waiting for bucket to be created, %v\n", bucketName)
		return err
	}

	fmt.Printf("Bucket %q successfully created\n", bucketName)
	return nil
}

// ListBuckets lists all buckets
func (svc *Service) ListBuckets() error {
	result, err := svc.Client.ListBuckets(nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to list buckets, %v\n", err)
		return err
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
	return nil
}

// ListBucketItems lists all the objects existing in the bucket with given name
func (svc *Service) ListBucketItems(bucketName string) error {
	result, err := svc.Client.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to list objects in the bucket with given name %s, %v\n", bucketName, err)
		return err
	}

	fmt.Printf("Object in the Bucket %s:\n", bucketName)

	for _, item := range result.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}
	return nil
}

// UploadFileToBucket lists all the objects existing in the bucket with given name
func (svc *Service) UploadFileToBucket(fileName string, bucketName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open file %s, %v\n", fileName, err)
		return err
	}
	defer file.Close()

	_, err = svc.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to upload file %s to bucket %s, %v\n", fileName, bucketName, err)
		return err
	}
	fmt.Printf("Successfully uploaded %s to %s\n", fileName, bucketName)
	return nil
}
