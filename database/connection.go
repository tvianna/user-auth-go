package database

import (
	"auth-go/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	databaseName := os.Getenv("DB_NAME")
	databasePass := os.Getenv("DB_ROOT_PASSWORD")
	databaseUser := os.Getenv("DB_USER")

	dsn := databaseUser + ":" + databasePass + "@tcp(" + host + ":" + port + ")/" + databaseName
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = connection

	connection.AutoMigrate(&models.User{})

}
