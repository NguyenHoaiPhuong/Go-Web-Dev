package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string `gorm:primary_key`
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	var product Product = Product{Code: "L1212", Price: 1000}

	// Create
	db.Create(&product)
	fmt.Println(product)

	product.Code = ""
	product.Price = 2000
	b := db.NewRecord(product)
	fmt.Println("b =", b)
	db.Create(&product)
	fmt.Println(product)

	// Read
	db.First(&product, 1) // find product with id 1
	fmt.Println(product)

	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	//db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)
}
