package config

import (
	"GO-WEB-DEV/101_Best-Practices/01_Practice/error"
	"GO-WEB-DEV/101_Best-Practices/01_Practice/jsonfunc"
)

// Config includes all initial settings
type Config struct {
	DatabaseConfig *DBConfig `json:"databaseConfig"`
}

// DBConfig includes all inital settings for database
type DBConfig struct {
	Dialect string `json:"dialect"`
	Source  string `json:"source"`
}

var cf *Config

// GetConfig initializes the config and returns it
func GetConfig(fileName string) (*Config, error.Error) {
	if cf == nil {
		cf = new(Config)
		err := jsonfunc.ReadFromFile(fileName, cf)
		if err != nil {
			var errNew error.ErrorImp
			errNew.InsertErrorMessage(err.Error())
			errNew.InsertErrorMessage(error.ErrorSetConfig)
			return nil, errNew
		}
	}
	return cf, nil
}
