package product

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"net/http"
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

	// 1. get all product records
	initializers.DB.Model(&models.Product{}).Find(&products)

	c.JSON(http.StatusOK, gin.H{"product": products})
}

func GetProductByID(c *gin.Context) {
	var product models.Product

	// 1. get the product record
	result := initializers.DB.Model(&models.Product{}).Where("id = ?", c.Param("id")).First(&product)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "document not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}
