package handlers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const aggregateFilePath = "./uploads/aggregate.txt"

// TODO: - VERIFICAR A POSSIBILIDADE DE ADICIONAR UM BANCO DE VETOR

func FeedDocuments(c *gin.Context) {
	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.Mkdir(uploadDir, os.ModePerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create directory"})
			return
		}
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file"})
		return
	}

	if filepath.Ext(file.Filename) != ".txt" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only .txt files are supported"})
		return
	}

	newFileName := time.Now().Format("20060102_150405") + ".txt"
	filePath := filepath.Join(uploadDir, newFileName)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	aggregateContent := ""
	if _, err := os.Stat(aggregateFilePath); err == nil {
		existingContent, err := ioutil.ReadFile(aggregateFilePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read aggregate document"})
			return
		}
		aggregateContent = string(existingContent)
	}

	aggregateContent += "\n\n" + string(content)

	err = ioutil.WriteFile(aggregateFilePath, []byte(aggregateContent), os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update aggregate document"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":           "File uploaded and added to aggregate document successfully",
		"file_name":         newFileName,
		"file_content":      string(content),
		"aggregate_updated": true,
	})
}
