package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// AppConfig holds all the configuration variables for the application
type AppConfig struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
	FrontendURL string
	Environment string
}

// Envs is the global configuration object
var Envs AppConfig

// LoadConfig initializes the configuration variables.
// It tries to load from a .env file first, then falls back to system environment variables.
func LoadConfig() {
	// 1. Attempt to load .env file
	// We ignore the error because on production (Render/Docker), there is no .env file
	_ = godotenv.Load()

	// 2. Populate the struct with values or defaults
	Envs = AppConfig{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		JWTSecret:   getEnv("JWT_SECRET", "default-secret-do-not-use-in-prod"),
		FrontendURL: getEnv("FRONTEND_URL", "http://localhost:3000"),
		Environment: getEnv("ENV", "development"),
	}

	// 3. Fail Fast: If critical variables are missing, stop the server immediately
	if Envs.DatabaseURL == "" {
		log.Fatal("CRITICAL ERROR: DATABASE_URL is not set in environment variables.")
	}
}

// getEnv retrieves the value of the environment variable named by the key.
// If the variable is present, its value is returned. Otherwise, the fallback is returned.
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
