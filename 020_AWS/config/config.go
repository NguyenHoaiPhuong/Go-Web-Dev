package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configurations exported
type Configurations struct {
	S3config S3Configurations `json:"s3config"`
}

// S3Configurations exported
type S3Configurations struct {
	Region string `json:"region"`
	Bucket string `json:"bucket"`
}

func (c *Configurations) init() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath("./resources")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("s3config.region", "us-east-1")
	viper.SetDefault("s3config.bucket", "temporary")

	err := viper.Unmarshal(c)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
}

// Print : prints all configurations on the terminal
func (c *Configurations) Print() {
	// Reading variables using the model
	fmt.Println("Reading variables using the model..")
	fmt.Println("Region is\t", c.S3config.Region)
	fmt.Println("Bucket is\t\t", c.S3config.Bucket)

	// Reading variables without using the model
	fmt.Println("\nReading variables without using the model..")
	fmt.Println("Region is\t", viper.GetString("s3config.region"))
	fmt.Println("Bucket is\t\t", viper.GetString("s3config.bucket"))
}

// NewConfigurations init new config and return it
func NewConfigurations() *Configurations {
	c := new(Configurations)
	c.init()

	return c
}
