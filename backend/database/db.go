package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Pool is the global database connection pool.
// By capitalizing "Pool", it becomes public and accessible from other packages.
var Pool *pgxpool.Pool

// Connect initializes the PostgreSQL connection pool using pgxpool.
// It takes the database URL as a parameter (which we will get from config).
func Connect(dbURL string) {
	var err error

	// 1. Create the connection pool
	Pool, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("CRITICAL: Unable to create database connection pool: %v\n", err)
	}

	// 2. Ping the database to ensure the connection is actually working
	err = Pool.Ping(context.Background())
	if err != nil {
		log.Fatalf("CRITICAL: Database ping failed: %v\n", err)
	}

	log.Println("SUCCESS: Connected to PostgreSQL Database!")
}

// Close gracefully closes all database connections in the pool.
// This should be called when the server is shutting down.
func Close() {
	if Pool != nil {
		Pool.Close()
		log.Println("Database connection pool closed.")
	}
}
