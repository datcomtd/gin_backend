package booking

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"time"
	"net/http"
	"strconv"
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

	// Get date range parameters
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	// Initialize DB query
	query := initializers.DB.Model(&models.Booking{})

	// Apply date range filter if provided
	if startDateStr != "" && endDateStr != "" {
		startDate, err1 := time.Parse("2006-01-02", startDateStr)
		endDate, err2 := time.Parse("2006-01-02", endDateStr)
		if err1 == nil && err2 == nil && !startDate.After(endDate) {
			query = query.Where("timestamp_start >= ? AND timestamp_end <= ?", startDate, endDate)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date range"})
			return
		}
	}

	// Get total count of matching records
	query.Count(&totalCount)

	// Fetch paginated results
	result := query.Limit(limit).Offset(offset).Find(&bookings)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_count": totalCount,
		"page":        page,
		"limit":       limit,
		"bookings":    bookings,
	})
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
