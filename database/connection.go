package database

import (
	"auth-go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	connection, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3200)/auth"), &gorm.Config{})

	if err != nil {
		panic("could not connect with database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})

}
