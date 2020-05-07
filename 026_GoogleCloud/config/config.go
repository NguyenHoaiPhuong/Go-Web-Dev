package config

import (
	"log"
	"os"
)

var config *Config

// Config struct
type Config struct {
	GoogleCredentials    string `json:"google_credentials"`
	GoogleBucketName     string `json:"google_bucket_name"`
	GoogleBaseStorageURL string `json:"google_base_storage_url"`
}

// GetConfig : getter
func GetConfig() *Config {
	return config
}

func init() {
	googleCredentials := os.Getenv("google_credentials")
	googleBucketName := os.Getenv("google_bucket_name")
	googleBaseStorageURL := os.Getenv("google_base_storage_url")

	log.Printf(`
	google_credentials: %s
	google_bucket_name: %s
	google_base_storage_url: %s
	`,
		googleCredentials,
		googleBucketName,
		googleBaseStorageURL,
	)

	config = &Config{
		GoogleCredentials:    googleCredentials,
		GoogleBucketName:     googleBucketName,
		GoogleBaseStorageURL: googleBaseStorageURL,
	}
}
