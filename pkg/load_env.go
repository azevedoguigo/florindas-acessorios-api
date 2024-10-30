package pkg

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found!")
		log.Println("Please, add .env file in root directory.")
	}
}
