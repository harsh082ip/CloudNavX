package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GoogleClientID     string
	GoogleClientSecret string
	GoogleCallbackURL  string
	ServerPort         string
	RedisURI           string
	SessionSecret      string
}

var AppConfig Config

func LoadConfig() {
	// Load .env
	// fmt.Println(os.Getwd())
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AppConfig = Config{
		GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
		GoogleCallbackURL:  getEnv("GOOGLE_CALLBACK_URL", "http://localhost:8000/auth/google/callback"),
		ServerPort:         getEnv("SERVER_PORT", ":8000"),
		RedisURI:           getEnv("REDIS_URI", ""),
		SessionSecret:      getEnv("SESSION_SECRET", ""),
	}

	fmt.Println("Loaded Configurations...")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
