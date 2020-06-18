package config

import (
	"log"
	"os"
)

var config *Config

func init() {
	// google
	googleCredentials := os.Getenv("google_credentials")
	googleBucketName := os.Getenv("google_bucket_name")
	googleBaseStorageURL := os.Getenv("google_base_storage_url")
	googleAPIKey := os.Getenv("google_api_key")

	// we dont use zap log here in order to see easier for development
	log.Printf(`	
	google_credentials: %s | google_bucket_name: %s | google_base_storage_url: %s | google_api_key: %s
	`,
		googleCredentials, googleBucketName, googleBaseStorageURL, googleAPIKey,
	)

	config = &Config{
		GoogleCredentials:    googleCredentials,
		GoogleBucketName:     googleBucketName,
		GoogleBaseStorageURL: googleBaseStorageURL,
		GoogleAPIKey:         googleAPIKey,
	}
}

// GetConfig :
func GetConfig() *Config {
	return config
}

// Config : struct
type Config struct {
	// google
	GoogleCredentials    string `json:"google_credentials"`
	GoogleBucketName     string `json:"google_bucket_name"`
	GoogleBaseStorageURL string `json:"google_base_storage_url"`
	GoogleAPIKey         string `json:"google_api_key"`
}
