package main

import (
	"log"
	"time"

	"github.com/Lezard82/movies-api/config"
	"github.com/Lezard82/movies-api/src/infrastructure/db"
	"github.com/Lezard82/movies-api/src/infrastructure/db/models"
	"gorm.io/gorm"
)

func SeedData(db *gorm.DB) {
	user := models.UserModel{Username: "admin", Password: "*3PMJq9729<y"}
	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("Error inserting test user: %v", err)
	}

	movie := models.MovieModel{
		Title:       "Inception",
		Director:    "Christopher Nolan",
		ReleaseDate: time.Date(2010, 7, 16, 0, 0, 0, 0, time.UTC),
		Cast:        `["Leonardo DiCaprio", "Joseph Gordon-Levitt"]`,
		Genre:       "Sci-Fi",
		Synopsis:    "A mind-bending thriller.",
		UserID:      user.ID,
	}

	if err := db.Create(&movie).Error; err != nil {
		log.Fatalf("Error inserting test movie: %v", err)
	}
}

func MigrateDB() {
	database := db.InitDB()
	if database == nil {
		log.Fatal("Failed to initialize database")
	}

	err := database.AutoMigrate(&models.UserModel{}, &models.MovieModel{})
	if err != nil {
		log.Fatalf("Error in AutoMigrate: %v", err)
	}

	SeedData(database)

	log.Println("Migration completed successfully!")
}

func main() {
	config.LoadEnv()
	MigrateDB()
}
