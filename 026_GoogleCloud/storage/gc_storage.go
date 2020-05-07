package storage

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/storage"
	"github.com/NguyenHoaiPhuong/Go-Web-Dev/026_GoogleCloud/config"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// GCStorage : connect to gcloud
// [Credentials] file path
// [BucketName] bucket's name
type GCStorage struct {
	Credentials string
	BucketName  string
}

// Init : create google cloud storage struct
//
// input [google storage config]
func Init(conf *config.Config) *GCStorage {
	return &GCStorage{
		Credentials: conf.GoogleCredentials,
		BucketName:  conf.GoogleBucketName,
	}
}

// InitClient : init google client
func (gc *GCStorage) InitClient() (*storage.Client, error) {
	ctx := context.Background()

	// Creates a client.
	opt := option.WithCredentialsFile(gc.Credentials)
	client, err := storage.NewClient(ctx, opt)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
		return nil, err
	}

	return client, nil
}

// List : list all objects in the given bucket
func List(client *storage.Client, bucket string) (objects []string, err error) {
	// [START storage_list_files]
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	it := client.Bucket(bucket).Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		objects = append(objects, attrs.Name)
	}
	// [END storage_list_files]
	return
}

// ListByPrefix :
func ListByPrefix(client *storage.Client, bucket, prefix, delim string) (objects []string, err error) {
	// [START storage_list_files_with_prefix]
	// Prefixes and delimiters can be used to emulate directory listings.
	// Prefixes can be used filter objects starting with prefix.
	// The delimiter argument can be used to restrict the results to only the
	// objects in the given "directory". Without the delimiter, the entire tree
	// under the prefix is returned.
	//
	// For example, given these blobs:
	//   /a/1.txt
	//   /a/b/2.txt
	//
	// If you just specify prefix="a/", you'll get back:
	//   /a/1.txt
	//   /a/b/2.txt
	//
	// However, if you specify prefix="a/" and delim="/", you'll get back:
	//   /a/1.txt
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	it := client.Bucket(bucket).Objects(ctx, &storage.Query{
		Prefix:    prefix,
		Delimiter: delim,
	})
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		objects = append(objects, attrs.Name)
	}
	// [END storage_list_files_with_prefix]
	return
}

// Delete : delete object in a bucket
func Delete(client *storage.Client, bucket, object string) error {
	// [START delete_file]
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()
	o := client.Bucket(bucket).Object(object)
	if err := o.Delete(ctx); err != nil {
		return err
	}
	// [END delete_file]
	return nil
}
