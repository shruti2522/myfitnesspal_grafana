package database

import (
	"os"
	"log"

    "fmt"
    "github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
    _ "github.com/jinzhu/gorm/dialects/postgres"
	"myfitnesspal-grafana/models"
)

func InitDB() (*gorm.DB, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASS")

	connectionString := fmt.Sprintf(" host=%s user=%s dbname=%s password=%s",host, user, dbname, password)

	db, err := gorm.Open("postgres", "sslmode=disable" + connectionString)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		os.Exit(1)
	}
	db.AutoMigrate(&models.UserEntry{})

    return db, nil
}
