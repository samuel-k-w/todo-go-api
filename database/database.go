package database

import (
	"log"
	"todo-go/internals/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	err = db.AutoMigrate(&models.Todo{}, &models.User{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	DB = db
	log.Println("connected to sqlite")
}
