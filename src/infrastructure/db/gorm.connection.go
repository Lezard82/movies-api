package db

import (
	"fmt"
	"log"
	"os"

	"github.com/Lezard82/movies-api/src/infrastructure/db/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB inicializa y devuelve una conexión a la base de datos
func InitDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting with database: %v", err)
	}

	err = database.AutoMigrate(&models.MovieModel{})
	if err != nil {
		log.Fatalf("Error in AutoMigrate: %v", err)
		return nil
	}

	return database
}
