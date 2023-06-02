package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func initConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func LoadConfig() string {
	envSource := "SYSTEM"

	if os.Getenv("BYPASS_ENV_FILE") == "" {
		initConfig()
		envSource = "FILE"
	}

	return envSource
}
