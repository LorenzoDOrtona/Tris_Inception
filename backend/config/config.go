package config

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port        string
	DatabaseURL string // This will be the final assembled string
	JWTSecret   string
	FrontendURL string
	Environment string
}

var Envs AppConfig

func LoadConfig() {
	_ = godotenv.Load()

	// 1. Get database components from environment (injected by the external infrastructure)
	dbUser := getEnv("DB_USER", "dbuser")
	dbPass := getEnv("DB_PASSWORD", "dbpassword")
	dbHost := getEnv("DB_HOST", "localhost")
	dbName := getEnv("DB_NAME", "myappdb")
	dbPort := getEnv("DB_PORT", "5432")

	// 2. Build the DatabaseURL string: postgres://user:password@host:port/dbname
	// We use url.QueryEscape for the password to handle special characters (@, #, etc.)
	// and add sslmode=disable for Kubernetes environments without SSL.
	dbPassEncoded := url.QueryEscape(dbPass)
	fullDatabaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassEncoded, dbHost, dbPort, dbName)

	Envs = AppConfig{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", fullDatabaseURL), // Use assembled URL as default
		JWTSecret:   getEnv("JWT_SECRET", "default-secret-do-not-use-in-prod"),
		FrontendURL: getEnv("FRONTEND_URL", "http://localhost:3000"),
		Environment: getEnv("ENV", "development"),
	}

	if Envs.DatabaseURL == "" {
		log.Fatal("CRITICAL ERROR: DATABASE_URL is not set.")
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
