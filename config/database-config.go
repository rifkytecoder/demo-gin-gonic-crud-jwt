package config

import (
	"fmt"
	"lab-go-gin-jwt/entity"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Creating a new connection to our database
func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("failed to load .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	// models in here
	db.AutoMigrate(&entity.Book{}, &entity.User{})
	return db
}

// Method is closing a connection between your app and your db
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
