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

// ListBuckets lists all buckets
func (svc *Service) ListBuckets() {
	result, err := svc.Client.ListBuckets(nil)
	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
	}

	fmt.Println("Buckets:")

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}

// CreateNewBucket creates new bucket
func (svc *Service) CreateNewBucket(bucketName string) {
	_, err := svc.Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		exitErrorf("Unable to create bucket %s, %v", bucketName, err)
	}

	// Wait until bucket is created before finishing
	fmt.Printf("Waiting for bucket %q to be created...\n", bucketName)

	err = svc.Client.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		exitErrorf("Error occurred while waiting for bucket to be created, %v", bucketName)
	}

	fmt.Printf("Bucket %q successfully created\n", bucketName)
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
