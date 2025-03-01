package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv carga las variables de entorno desde el archivo .env
func LoadEnv() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Println("Warning: The .env file could not be loaded; system environment variables will be used instead.")
	}
}
