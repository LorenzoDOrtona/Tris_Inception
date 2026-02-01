package main

import (
	"net/http"

	"github.com/LorenzoDOrtona/Tris_Inception/cmd/server/api"
	internal "github.com/LorenzoDOrtona/Tris_Inception/internal/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setup_gin(GC *internal.GameController) {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()
	// CONFIGURAZIONE CORS: Permetti al frontend di chiamare il backend
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"https://trisinception.onrender.com"} // L'indirizzo di React
	config.AllowMethods = []string{"GET", "POST", "PUT", "OPTIONS"}
	r.Use(cors.New(config))
	r.POST("/playerToken", func(c *gin.Context) {
		var RQ api.ReqToken
		if err := c.ShouldBindJSON(&RQ); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		response := GC.CreatePlayerToken(&RQ)
		c.JSON(200, response)
	})

	r.POST("/match", func(c *gin.Context) {
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
	r.POST("/move", func(c *gin.Context) {
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
	r.POST("/polling", func(c *gin.Context) {
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
