package config

import (
	"GO-WEB-DEV/080_Json/jsonfunc"
)

var cf *Config

// GetConfig returns global var cf
func GetConfig(fileName string) (*Config, error) {
	if cf == nil {
		cf = new(Config)
		err := jsonfunc.ReadFromFile(fileName, cf)
		if err != nil {
			return nil, err
		}
	}
	return cf, nil
}

// Config includes initial settings
type Config struct {
	*S3Config `json:"s3config"`
}

// S3Config includes AWS S3 settings
type S3Config struct {
	Region string `json:"region"`
	Bucket string `json:"bucket"`
}
