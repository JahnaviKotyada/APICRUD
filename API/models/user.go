package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=jahnavi@2003 dbname=info port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	database.AutoMigrate(&User{})
	DB = database
}
