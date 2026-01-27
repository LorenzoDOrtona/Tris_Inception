package main

import (
	"net/http"

	"github.com/LorenzoDOrtona/Tris_Inception/cmd/server/api"
	internal "github.com/LorenzoDOrtona/Tris_Inception/internal/controller"
	"github.com/gin-gonic/gin"
)

func setup_gin(GC *internal.GameController) {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()
	r.POST("/playerToken", func(c *gin.Context) {
		var RQ api.ReqToken
		if err := c.ShouldBindJSON(&RQ); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		response := GC.CreatePlayerToken(&RQ)
		c.JSON(200, gin.H{"response": response})
	})
	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}

func main() {
	var gc = internal.NewGameController()
	setup_gin(gc)

}
