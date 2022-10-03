package db

import (
	"fmt"
	"golang-swagger/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// connect to sqlite

// connect to postgres
const (
	host   = "localhost"
	port   = "5432"
	user   = "postgres"
	pass   = "postgres"
	dbname = "hacktiv8-latihan-8"
)

func ConnectGorm() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname,
	)

	// connect
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	db.AutoMigrate(&models.People{})

	// Create
	// db.Create(&models.People{Name: "jaka", Address: "Perum Griya Rajeg Lestari"})

	// // Read
	// var people models.People
	// db.First(&people, 1)                  // find product with integer primary key
	// db.First(&people, "name = ?", "jaka") // find people with code jaka

	// // Update - update product's price to 200
	// db.Model(&people).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&people).Updates(models.People{Name: "jaka update", Address: "Perum griya update"}) // non-zero fields
	// db.Model(&people).Updates(map[string]interface{}{"Name": 200, "Address": "Perum griya update"})

	// // Delete - delete people
	// db.Delete(&people, 1)

	// db.Debug().AutoMigrate(models.People{})

	return db
}
