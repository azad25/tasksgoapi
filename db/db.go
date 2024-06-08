package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"taskmanager/models"
)

var DB *gorm.DB

func Connect() {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"))

	//connect database using gorm
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to Postgres", err)
	}
}

func Migrate() {
	DB.AutoMigrate(&models.Task{})
}
