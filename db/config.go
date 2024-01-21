package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is the global variable for the database connection
var DB *gorm.DB

// InitializeDB initializes the database connection
func InitializeDB() {
	dsn := "user:password@tcp(localhost:3306)/database?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Set DB to use a singular table name (e.g., User instead of Users)
	db = db.Debug().Table("user")

	DB = db
}

// CloseDB closes the database connection
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}

// AutoMigrateModels migrates the defined models to the database
func AutoMigrateModels() {
	//TODO: Add code here to auto-migrate your models to the database
	// err := DB.AutoMigrate(&YourModel{}) // Replace YourModel with your actual Gorm model
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
