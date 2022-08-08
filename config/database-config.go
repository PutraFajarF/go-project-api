package config

import (
	"fmt"
	"os"

	"github.com/PutraFajarF/go-project-api/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// SetupDatabaseConnection is creating a new connection to our database
func SetupDatabaseConnection() *gorm.DB {
	// If use localhost uncomment this section
	// errEnv := godotenv.Load()
	// if errEnv != nil {
	// 	panic("Failed to load env file")
	// }

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, dbUser, dbName, dbPass, dbPort)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		panic("Failed to create a connection to your database")
	}

	db = conn
	db.AutoMigrate(&entity.Book{}, &entity.User{})
	return db
}

// CloseDatabaseConnection is closing a connection between your app and your db
func CloseDatabaseConnection(db *gorm.DB) {
	dbPostgre := db.DB()
	err := dbPostgre.Close()
	if err != nil {
		panic("Failed to close connection from database")
	}
}
