package dashboard

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/authentication/token"
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"net/http"
)

type CountsResponse struct {
	Users     int64 `json:"users"`
	Bookings  int64 `json:"bookings"`
	Documents int64 `json:"documents"`
	Products  int64 `json:"products"`
}

// GetCounts returns the total count of users, bookings, documents, and products
func GetCounts(c *gin.Context) {
	var counts CountsResponse

	username, _, _, errCode, errString := token.VerifyToken(c.GetHeader("Authorization"))
	if username == "" {
		c.JSON(errCode, gin.H{"message": errString})
		return
	}

	// Count users
	if err := initializers.DB.Model(&models.User{}).Count(&counts.Users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count users"})
		return
	}

	// Count bookings
	if err := initializers.DB.Model(&models.Booking{}).Count(&counts.Bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count bookings"})
		return
	}

	// Count documents
	if err := initializers.DB.Model(&models.Document{}).Count(&counts.Documents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count documents"})
		return
	}

	// Count products
	if err := initializers.DB.Model(&models.Product{}).Count(&counts.Products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count products"})
		return
	}

	c.JSON(http.StatusOK, counts)
}