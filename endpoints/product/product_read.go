package product

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"net/http"
	"strconv"
)

//
// GetProducts (GET)
//  1. get all product records
//
// GetProductByID (GET :id)
//  1. get the product record
//

func GetProducts(c *gin.Context) {
	var products []models.Product
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
	query := initializers.DB.Model(&models.Product{})

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
	result := query.Limit(limit).Offset(offset).Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_count": totalCount,
		"page":        page,
		"limit":       limit,
		"products":    products,
	})
}

func GetProductByID(c *gin.Context) {
	var product models.Product

	// 1. get the product record
	result := initializers.DB.Model(&models.Product{}).Where("id = ?", c.Param("id")).First(&product)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func GetProductByCategory(c *gin.Context) {
	var products []models.Product

	result := initializers.DB.Model(&models.Product{}).Where("category LIKE ?", c.Param("category")).Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count":   result.RowsAffected,
		"product": products})
}
