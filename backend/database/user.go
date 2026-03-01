package database

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	// Import your database package where Pool is defined
	// "your_module_name/database"
)

// CreateUsersTable initializes the users table in the PostgreSQL database.
// It uses password_hash instead of password to clearly indicate we store hashes.
func CreateUsersTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password_hash VARCHAR(255) NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);
	`

	// Execute the query using the global database Pool
	_, err := Pool.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	log.Println("SUCCESS: Users table is ready.")
	return nil
}

// hashPassword securely hashes a plain text password using bcrypt.
func hashPassword(password string) (string, error) {
	// bcrypt.DefaultCost is typically 10, balancing security and server CPU load
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedBytes), nil
}

// RegisterUser takes user details, hashes the password, and saves the record.
func RegisterUser(username string, email string, plainPassword string) error {
	// 1. Hash the plain password BEFORE talking to the database
	hashedPassword, err := hashPassword(plainPassword)
	if err != nil {
		return err
	}

	// 2. Insert the user with the hashed password, NOT the plain one
	query := `
		INSERT INTO users (username, email, password_hash) 
		VALUES ($1, $2, $3)
	`

	_, err = Pool.Exec(context.Background(), query, username, email, hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to insert user into database: %w", err)
	}

	log.Printf("SUCCESS: User '%s' registered securely.\n", username)
	return nil
}
