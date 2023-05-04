package configs

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	if godotenv.Load() != nil {
		log.Fatalln("Unable to load environment variables")
	}
}
