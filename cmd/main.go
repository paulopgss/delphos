package main

import (
	"delphos/config"
	"delphos/internal/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	router := gin.Default()

	router.GET("/health", handlers.HealthCheck)
	router.POST("/prompt", handlers.HandlePrompt)
	router.POST("/feed", handlers.FeedDocuments)
	router.POST("/trainModel", handlers.TrainModel)

	log.Printf("Server running on port %s", cfg.ServerPort)
	router.Run(":" + cfg.ServerPort)
}
