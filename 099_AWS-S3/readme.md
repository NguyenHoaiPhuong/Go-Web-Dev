Configuring the AWS SDK for Go
- Configure settings for service clients. Most settings are optional.
- Must specify a region and credentials. The SDK uses these values to send requests to the correct AWS Region and sign requests with the correct credentials
- You can specify these values as part of a session or as environment variables.

Create a Bucket:
- CreateNewBucket creates a new bucket with a given name

Delete a Bucket:
- The DeleteBucket function deletes a bucket.

List Buckets in S3:
- The ListBuckets lists all buckets in S3

List Bucket Items
- The ListBucketItems lists all items in a bucket

Upload a file to a Bucket
- The UploadFileToBucket uploads a file to a bucket

Download a file from a bucket
- The DownloadFileFromBucket downloads a file from a bucket

Delete an Item in a Bucket
- The DeleteObject function deletes an object from a bucket.

Delete all Items in a Bucket
- The DeleteObjects function deletes all objects from a bucket.