package booking

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"time"
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

	date := c.Param("day")
	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid date format"})
		return
	}
	endDate := startDate.AddDate(0, 0, 1) // Add one day to get the end of the day
	// 1. get the bookings record
	result := initializers.DB.Model(&models.Booking{}).
	Where("timestamp_start >= ? AND timestamp_start < ?", startDate, endDate).
	Find(&bookings)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "no booking found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count":   result.RowsAffected,
		"booking": bookings})
}

func ViewById(c *gin.Context) {
	var bookings []models.Booking

	id := c.Param("id")
	
	// 1. get the bookings record
	result := initializers.DB.Model(&models.Booking{}).Where("id = ?", id).
	Find(&bookings)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "no booking found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"count":   result.RowsAffected,
		"booking": bookings})
}
