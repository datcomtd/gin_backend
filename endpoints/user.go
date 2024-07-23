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
//  2. make the public_user record of the user
//

type public_User struct {
	Username string `json:"name"`
	Email    string `json:"email"`

	Role   uint `json:"role"`
	Course uint `json:"course"`
}

func GetUsers(c *gin.Context) {
	var users []public_User

	// 1. get all user records
	initializers.DB.Model(&models.User{}).Find(&users)

	c.JSON(http.StatusOK, gin.H{"user": users})
}

func GetUserByUsername(c *gin.Context) {
	var user public_User

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
