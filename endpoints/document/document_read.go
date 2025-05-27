package document

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"net/http"
	"strings"
	"strconv"
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
	var totalCount int64

	// Get pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	// Get filter parameters
	title := c.Query("title")
	category := c.Query("category")

	// Initialize DB query
	query := initializers.DB.Model(&models.Document{})

	// Apply title filter (partial match)
	if title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}

	// Apply category filter (exact match)
	if category != "" {
		query = query.Where("category = ?", category)
	}

	// Get total count of matching records
	query.Count(&totalCount)

	// Fetch paginated results
	result := query.Limit(limit).Offset(offset).Find(&documents)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_count": totalCount,
		"page":        page,
		"limit":       limit,
		"documents":   documents,
	})
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
	result := initializers.DB.Model(&models.Document{}).Where("category LIKE ?", strings.ToLower(c.Param("category"))).Find(&documents)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "document not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count":    result.RowsAffected,
		"document": documents})
}
