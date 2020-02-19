package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config : configurations
type Config struct {
	Server
	Database
}

// Server Configurations exported
type Server struct {
	Host string
	Port string
}

// Database Configurations exported
type Database struct {
	DBName string `json: "dbName"`
	DBType string `json: "dbType"`
	DBUser string `json: "dbUser"`
	DBPass string `json: "dbPass"`
	DBHost string `json: "dbHost"`
	DBPort string `json: "dbPort"`
}

func ParseConfig() *Config {
	// Set the file name of the configurations file
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	// Set the path to look for the configurations file
	viper.AddConfigPath("./resources")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	var conf Config

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set default db variables
	viper.SetDefault("database.dbName", "test_db")
	viper.SetDefault("database.dbType", "postgres")
	viper.SetDefault("database.dbUser", "postgres")
	viper.SetDefault("database.dbPass", "")
	viper.SetDefault("database.dbHost", "localhost")
	viper.SetDefault("database.dbPort", "5432")

	// Set default server variables
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", "3000")

	err := viper.Unmarshal(&conf)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading variables using the model
	fmt.Println("Database name is\t", conf.Database.DBName)
	fmt.Println("Database type is\t", conf.Database.DBType)

	return &conf
}
