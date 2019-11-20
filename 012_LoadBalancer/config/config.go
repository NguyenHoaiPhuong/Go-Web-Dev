package config

import (
	"github.com/spf13/viper"
)

// ReadConfig func
func ReadConfig() (Proxy, error) {
	proxy := Proxy{}

	// Set the file name of the configurations file
	viper.SetConfigName("config")
	// Set the path to look for the configurations file
	viper.AddConfigPath("./resources")
	// Set type
	viper.SetConfigType("yml")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return proxy, err
	}

	// Set default values
	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", "3000")
	viper.SetDefault("scheme", "http")

	err := viper.Unmarshal(&proxy)
	if err != nil {
		return proxy, err
	}

	return proxy, nil
}
