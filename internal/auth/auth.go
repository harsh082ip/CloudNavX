package auth

import (
	"log"

	"github.com/harsh082ip/CloudNavX/internal/config"
	redis_db "github.com/harsh082ip/CloudNavX/internal/db/redis"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func SetupGoogleAuth() {
	// Fetch configuration values
	ClientID := config.AppConfig.GoogleClientID
	ClientSecret := config.AppConfig.GoogleClientSecret
	CallbackURL := config.AppConfig.GoogleCallbackURL

	// Ensure required configuration values are present
	if ClientID == "" || ClientSecret == "" || CallbackURL == "" {
		log.Fatal("Google OAuth credentials or session secret are missing!")
	}

	// Initialize Redis store for session storage
	redisStore := redis_db.NewRedisStore()
	gothic.Store = redisStore

	// Setup Google provider
	goth.UseProviders(
		google.New(ClientID, ClientSecret, CallbackURL, "email"),
	)

	log.Println("Google OAuth with Redis session store setup completed successfully.")
}
