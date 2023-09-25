package main

import (
	"blogging/middleware"
	"blogging/routes"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
func main() {
	r := routes.SetupRoutes()
	// Database setup and migrations
	db := setupDB()

	// Use the authentication middleware for protected routes
	r.Use(middleware.Authenticate())

	// Use the error handling middleware
	r.Use(middleware.ErrorHandler())
	// Initialize the routes

	// Routes
	routes.SetupRoutes(r, db)

	r.Run(":8080")
}

func setupDB() *gorm.DB {
	// Initialize and return your PostgreSQL database connection here
}

func migrateDB(db *gorm.DB) {
	// Perform database migrations using gorm here
}
