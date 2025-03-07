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
	result := initializers.DB.Model(&models.Product{}).Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"count":   result.RowsAffected,
		"product": products})
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
