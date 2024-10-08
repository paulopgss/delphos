package handlers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

const trainFilePath = "./uploads/aggregate.txt"

func TrainModel(c *gin.Context) {
	var request struct {
		Question string `json:"question" binding:"required"`
		Answer   string `json:"answer" binding:"required"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request, both question and answer are required"})
		return
	}

	newEntry := "\n\nQ: " + request.Question + "\nA: " + request.Answer

	aggregateContent := ""
	if _, err := os.Stat(trainFilePath); err == nil {
		existingContent, err := ioutil.ReadFile(trainFilePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read aggregate document"})
			return
		}
		aggregateContent = string(existingContent)
	}

	aggregateContent += newEntry

	err := ioutil.WriteFile(trainFilePath, []byte(aggregateContent), os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update aggregate document"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Training data added successfully",
		"question": request.Question,
		"answer":   request.Answer,
	})
}
