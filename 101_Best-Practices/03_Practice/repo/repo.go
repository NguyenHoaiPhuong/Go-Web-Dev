package repo

import (
	"fmt"

	"github.com/NguyenHoaiPhuong/Go-Web-Dev/101_Best-Practices/03_Practice/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// Init func
func Init(conf config.Database) {
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		conf.DBHost, conf.DBPort, conf.DBUser, conf.DBName, conf.DBPass)
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	// db.Debug().AutoMigrate(&models.Account{}, &models.Contact{})
}

// GetDB func
func GetDB() *gorm.DB {
	return db
}
