package handlers

import (
	"delphos/internal/models"
	"delphos/internal/services"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func HandlePrompt(c *gin.Context) {
	var request models.PromptRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	context := ""
	if _, err := os.Stat(aggregateFilePath); err == nil {
		content, err := ioutil.ReadFile(aggregateFilePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read aggregate document"})
			return
		}
		context = string(content)
	}

	fullPrompt := "Context: " + context + "\n\nUser prompt: " + request.Prompt

	responseStream, err := services.SendPromptToModelStream(fullPrompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Model error"})
		return
	}

	var fullResponse strings.Builder
	for part := range responseStream {
		fullResponse.WriteString(part)
	}

	c.JSON(http.StatusOK, gin.H{"response": fullResponse.String()})
}
