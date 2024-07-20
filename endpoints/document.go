package endpoints

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"net/http"
	"strings"
)

//
// GetDocuments(GET)
//  1. get all document records
//
// GetDocumentByID (GET :id)
//  1. get the document record
//
// GetDocumentsByCategory (GET :category)
//  1. get all document records by category
//

func GetDocuments(c *gin.Context) {
	var documents []models.Document

	// 1. get all document records
	initializers.DB.Model(&models.Document{}).Find(&documents)

	c.JSON(http.StatusOK, gin.H{"document": documents})
}

func GetDocumentByID(c *gin.Context) {
	var document models.Document

	// 1. get the document record
	result := initializers.DB.Model(&models.Document{}).Where("id = ?", c.Param("id")).First(&document)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "document not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"document": document})
}

func GetDocumentsByCategory(c *gin.Context) {
	var documents []models.Document

	// 1. get all document records by category
	initializers.DB.Model(&models.Document{}).Where("category = ?", strings.ToLower(c.Param("category"))).Find(&documents)

	c.JSON(http.StatusOK, gin.H{"document": documents})
}
