package endpoints

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"net/http"
)

//
// GetMembers (GET)
//  1. get all user records
//
// GetUserByUsername (GET :username)
//  1. get the user record
//

func GetUsers(c *gin.Context) {
	var users []models.User

	// 1. get all user records
	initializers.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"user": users})
}

func GetUserByUsername(c *gin.Context) {
	var user models.User

	// 1. get the user record
	// 1.1. set the primaryKey to null
	user.Username = ""
	// 1.2. sql query
	result := initializers.DB.Model(&models.User{}).Where("username = ?", c.Param("username")).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
