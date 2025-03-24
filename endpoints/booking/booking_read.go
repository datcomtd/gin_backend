package booking

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"net/http"
)

//
// View (GET)
//  1. get all product records
//
// ViewByDay (GET :day)
//  1. get the product record
//

func View(c *gin.Context) {
	var bookings []models.Booking

	// 1. get all booking records
	result := initializers.DB.Model(&models.Booking{}).Find(&bookings)

	c.JSON(http.StatusOK, gin.H{
		"count":   result.RowsAffected,
		"booking": bookings})
}

func ViewByDay(c *gin.Context) {
	var bookings []models.Booking

	// 1. get the bookings record
	result := initializers.DB.Model(&models.Booking{}).Where("day = ?", c.Param("day")).Find(&bookings)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "no booking found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count":   result.RowsAffected,
		"booking": bookings})
}
