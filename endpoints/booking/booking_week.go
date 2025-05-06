package booking

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/authentication/token"
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"net/http"
	"time"
)
// ultimos dez documentos em ingles  

func GetCurrentWeekBookings(c *gin.Context) {
	var bookings []models.Booking

	username, _, _, errCode, errString := token.VerifyToken(c.GetHeader("Authorization"))
	if username == "" {
		c.JSON(errCode, gin.H{"message": errString})
		return
	}
	// Get current time
	now := time.Now()

	// Calculate start of the week (Monday)
	weekday := now.Weekday()
	daysSinceMonday := (int(weekday) + 6) % 7 // Convert Sunday=0 to Sunday=6
	startOfWeek := now.AddDate(0, 0, -daysSinceMonday).Truncate(24 * time.Hour)

	// Calculate end of the week (Sunday)
	endOfWeek := startOfWeek.AddDate(0, 0, 7).Add(-time.Nanosecond)


	result := initializers.DB.Model(&models.Booking{}).Where("timestamp_start >= ? AND timestamp_start <= ?", startOfWeek, endOfWeek).Find(&bookings);
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "no booking found"})
		return
	}

	c.JSON(http.StatusOK, bookings)

}