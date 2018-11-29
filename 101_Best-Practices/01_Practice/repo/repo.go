package repo

import (
	"GO-WEB-DEV/101_Best-Practices/01_Practice/config"
	"GO-WEB-DEV/101_Best-Practices/01_Practice/error"
	"GO-WEB-DEV/101_Best-Practices/01_Practice/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var database *gorm.DB

// GetDatabase returns database
func GetDatabase() (*gorm.DB, error.Error) {
	if database == nil {
		err := initDatabase()
		if err != nil {
			return nil, err
		}
	}
	return database, nil
}

// initDatabase initializes database
func initDatabase() error.Error {
	cf, err := config.GetConfig("resource/config.json")
	if err != nil {
		var errNew error.ErrorImp
		errNew.InsertErrorMessage(err.Error())
		errNew.InsertErrorMessage(error.ErrorDatabaseConnection)
		return errNew
	}
	dbConfig := cf.DatabaseConfig
	db, osErr := gorm.Open(dbConfig.Dialect, dbConfig.Source)
	if osErr != nil {
		var errNew error.ErrorImp
		errNew.InsertErrorMessage(osErr.Error())
		errNew.InsertErrorMessage(error.ErrorDatabaseConnection)
		return errNew
	}
	database = db.AutoMigrate(&model.Employee{})
	return nil
}
