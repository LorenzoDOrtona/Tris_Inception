package main

import (
	"log"
	"net/http"

	"github.com/LorenzoDOrtona/Tris_Inception/cmd/server/api"
	"github.com/LorenzoDOrtona/Tris_Inception/config"
	"github.com/LorenzoDOrtona/Tris_Inception/database"
	internal "github.com/LorenzoDOrtona/Tris_Inception/internal/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setup_gin(GC *internal.GameController) {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()
	// CORS CONFIGURATION: Allow the frontend to call the backend
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"https://tris.lorenzodortona.com",
		"http://localhost",
		"http://localhost:5173",
	}
	config.AllowMethods = []string{"GET", "POST", "PUT", "OPTIONS"}
	r.Use(cors.New(config))
	apiR := r.Group("/api")
	{

		apiR.POST("/match", func(c *gin.Context) {
			var reqGame api.ReqCreateOrJoinGame
			// Bind JSON input to struct
			if err := c.ShouldBindJSON(&reqGame); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			//check game logic success
			response, err := GC.CreateGame(&reqGame)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			c.JSON(200, response)

		})
		apiR.POST("/move", func(c *gin.Context) {
			var reqMove api.ReqMove
			// Bind JSON input to struct
			if err := c.ShouldBindJSON(&reqMove); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			//check game logic success
			response, err := GC.MakeMove(&reqMove)
			if err != nil {
				c.JSON(400, api.RespError{ErrorMessage: err.Error()})
				return
			}
			c.JSON(200, response)

		})
		apiR.POST("/polling", func(c *gin.Context) {
			var reqPool api.ReqPooling
			// Bind JSON input to struct
			if err := c.ShouldBindJSON(&reqPool); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			//check game logic success
			response, err := GC.CheckLastTimestamp(&reqPool)
			if err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}
			c.JSON(200, response)

		})
		// Define a simple GET endpoint
		apiR.GET("/ping", func(c *gin.Context) {
			// Return JSON response
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		apiR.POST("/register", func(c *gin.Context) {
			var req api.ReqRegister
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(400, gin.H{"error": "Invalid input: " + err.Error()})
				return
			}

			// lets call the db function
			err := database.RegisterUser(req.Username, req.Email, req.Password)
			if err != nil {
				c.JSON(500, gin.H{"error": "Registration failed: " + err.Error()})
				return
			}

			c.JSON(201, gin.H{"message": "User created successfully"})
		})
		apiR.POST("/login", func(c *gin.Context) {
			var req api.ReqToken // You'll need to define this struct in your api package
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(400, gin.H{"error": "Invalid input"})
				return
			}

			// Verification logic (to be implemented in your database package)
			user, err := GC.CreatePlayerToken(&req)
			if err != nil {
				c.JSON(401, gin.H{"error": "Invalid credentials"})
				return
			}

			c.JSON(200, gin.H{
				"message":  "Login successful",
				"username": req.PlayerName,
				"token":    user.Token,
			})
		})
		apiR.POST("/guest", func(c *gin.Context) {
			user, err := GC.CreateGuestToken()
			if err != nil {
				c.JSON(500, gin.H{"error": "Failed to create guest token"})
				return
			}

			c.JSON(200, gin.H{
				"message":  "Guest login successful",
				"username": "Guest",
				"token":    user.Token,
			})
		})
	}
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}

func main() {
	// 1. Load Configuration
	config.LoadConfig()

	// 2. Connect to Database using the URL we just built
	database.Connect(config.Envs.DatabaseURL)
	defer database.Close()

	// 3. Setup Tables
	// Note: Move CreateUsersTable to a separate package like 'repository'
	// to avoid circular dependencies with the 'database' package.
	err := database.CreateUsersTable()
	if err != nil {
		log.Fatalf("Could not setup database tables: %v", err)
	}

	log.Printf("Server starting on port %s in %s mode", config.Envs.Port, config.Envs.Environment)

	// 4. Starting HTTP server here...
	var gc = internal.NewGameController()
	setup_gin(gc)

}
