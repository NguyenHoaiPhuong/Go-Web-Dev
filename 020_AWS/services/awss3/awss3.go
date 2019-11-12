package awss3

import (
	"fmt"
	"os"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/020_AWS/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

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

// Service implements all methods declared in the interface
type Service struct {
	IService

	Session    *session.Session
	Client     *s3.S3
	Uploader   *s3manager.Uploader
	Downloader *s3manager.Downloader
}

// Init initializes settings
func (svc *Service) init(s3config config.S3Configurations) {
	config := &aws.Config{
		Region: aws.String(s3config.Region),
	}
	svc.Session = session.Must(session.NewSession(config))
	/*val, err := svc.Session.Config.Credentials.Get()
	if err != nil {
		panic(err)
	}
	// print credentials
	fmt.Println(val)*/
	svc.Client = s3.New(svc.Session)
	svc.Uploader = s3manager.NewUploader(svc.Session)
	svc.Downloader = s3manager.NewDownloader(svc.Session)
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

// DeleteBucket deletes a bucket in S3
func (svc *Service) DeleteBucket(bucketName string) error {
	_, err := svc.Client.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete bucket %s, %v", bucketName, err)
		return err
	}

	// Wait until bucket is deleted before finishing
	fmt.Printf("Waiting for bucket %s to be deleted...\n", bucketName)

	err = svc.Client.WaitUntilBucketNotExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred while waiting for bucket to be deleted, %v\n", bucketName)
		return err
	}
	fmt.Printf("Bucket %s was deleted\n", bucketName)
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

// DeleteBucketItem deletes one item in a bucket
func (svc *Service) DeleteBucketItem(fileName string, bucketName string) error {
	_, err := svc.Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete object %s in the bucket %s, %v\n", fileName, bucketName, err)
		return err
	}
	err = svc.Client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred while waiting for object %s to be deleted, %v\n", fileName, err)
		return err
	}
	fmt.Printf("Deleted object %s from bucket %s\n", fileName, bucketName)
	return nil
}

// CopyItemFromBucketToBucket : copies item from bucket to bucket
func (svc *Service) CopyItemFromBucketToBucket(from string, to string, file string) error {
	source := from + "/" + file

	// Copy the item
	_, err := svc.Client.CopyObject(&s3.CopyObjectInput{Bucket: aws.String(to), CopySource: aws.String(source), Key: aws.String(file)})
	if err != nil {
		fmt.Printf("Unable to copy item from bucket %s to bucket %s, %v\n", from, to, err)
		return err
	}

	// Wait to see if the item got copied
	err = svc.Client.WaitUntilObjectExists(&s3.HeadObjectInput{Bucket: aws.String(to), Key: aws.String(file)})
	if err != nil {
		fmt.Printf("Error occurred while waiting for item %s to be copied to bucket %q, %v\n", file, to, err)
		return err
	}

	fmt.Printf("Item %s successfully copied from bucket %s to bucket %s\n", file, from, to)
	return nil
}

// DeleteAllBucketItems deletes all items in a bucket
func (svc *Service) DeleteAllBucketItems(bucketName string) error {
	iter := s3manager.NewDeleteListIterator(svc.Client, &s3.ListObjectsInput{
		Bucket: aws.String(bucketName),
	})
	err := s3manager.NewBatchDeleteWithClient(svc.Client).Delete(aws.BackgroundContext(), iter)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to delete objects from bucket %s, %v\n", bucketName, err)
		return err
	}
	fmt.Printf("Deleted object(s) from bucket: %s", bucketName)
	return nil
}

// UploadFileToBucket uploads a file to bucket
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

// DownloadFileFromBucket downloads a file from a bucket
func (svc *Service) DownloadFileFromBucket(fileName string, bucketName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create file %s, %v\n", fileName, err)
		return err
	}
	numBytes, err := svc.Downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to download item %s from bucket %s, %v", fileName, bucketName, err)
		return err
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
	return nil
}

// NewService returns new awss3 service
func NewService(s3config config.S3Configurations) IService {
	svc := new(Service)
	svc.init(s3config)

	return svc
}
