package user

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"net/http"
	"strconv"
)

//
// GetMembers (GET)
//  1. get all user records
//
// GetUserByUsername (GET :username)
//  1. get the user record
//  2. make the public_user record of the user
//

type public_User struct {
	Username string `json:"name"`
	Email    string `json:"email"`

	Role   uint `json:"role"`
	Course uint `json:"course"`
	RA     string `json:"ra"`
}

type public_User_profile struct {
	Username string `json:"name"`
	Email    string `json:"email"`

	Role   		uint `json:"role"`
	Course 		uint `json:"course"`
	RA     		string `json:"ra"`
	Picture 	string `json:"picture"`
	CreatedAt 	string `json:"created_at"`
}

func GetUsers(c *gin.Context) {
	var users []models.User
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
	username := c.Query("username")
	courseStr := c.Query("course")

	// Initialize DB query
	query := initializers.DB.Model(&models.User{})

	// Apply username filter (partial match)
	if username != "" {
		query = query.Where("username ILIKE ?", "%"+username+"%")
	}

	// Apply course filter (exact match)
	if courseStr != "" {
		course, err := strconv.Atoi(courseStr)
		if err == nil && course >= 1 && course <= 8 { // Validate against Course constants (1-8)
			query = query.Where("course = ?", course)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course value"})
			return
		}
	}

	// Get total count of matching records
	query.Count(&totalCount)

	// Fetch paginated results
	result := query.Limit(limit).Offset(offset).Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_count": totalCount,
		"page":        page,
		"limit":       limit,
		"users":       users,
	})
}

func GetUserByUsername(c *gin.Context) {
	var user models.User

	// 1. get the user record
	username := c.Param("username")
	// 1.1. sql query
	result := initializers.DB.Model(&models.User{}).Where("username = ?", username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	// 1.2. check if the username is valid
	if username != user.Username {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
