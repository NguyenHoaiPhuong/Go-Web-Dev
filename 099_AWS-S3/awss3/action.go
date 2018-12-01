package awss3

import (
	"GO-WEB-DEV/099_AWS-S3/config"
)

// Actions includes all functions related to S3
type Actions interface {
	Init(cf *config.S3Config)
	ListBuckets()
}
