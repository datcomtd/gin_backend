package endpoints

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/authentication"
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"net/http"
	"strings"
)

//
// UpdateUser (POST updateRequest)
//  0. retrieve post data
//  1. check if the username and password is filled
//  2. check if the user exists
//  3. check if the password is correct
//  4. update the user struct fields
//  5. update the user record in the database
//

type user_updateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`

	NewPassword string `json:"newpassword"`
	Email       string `json:"email"`

	Role   uint `json:"role"`
	Course uint `json:"course"`
}

func UpdateUser(c *gin.Context) {
	var body user_updateRequest
	var user models.User

	//  0. retrieve post data
	c.Bind(&body)

	// 1. check if the username and password is filled
	if (body.Username == "") || (body.Password == "") {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required fields are not filled"})
		return
	}

	// 2. check if the user exists
	result := initializers.DB.Where("username = ?", strings.ToLower(body.Username)).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid username or password"})
		return
	}

	// 3. check if the password is correct
	bl := authentication.VerifyPassword(body.Password, user.Password)
	if bl != true {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid username or password"})
		return
	}

	// 4. update the user struct fields
	// 4.1. new password field
	if body.NewPassword != "" {
		// 4.1.1. hash the provided new password string
		hashedPassword, err := authentication.HashPassword(body.NewPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed hashing the password"})
			return
		}
		// 4.1.2. update the new password field
		user.Password = hashedPassword
	}
	// 4.2. email field
	if body.Email != "" {
		user.Email = body.Email
	}
	// 4.3. role field
	if body.Role != 0 {
		user.Role = body.Role
	}
	// 4.4. course field
	if body.Course != 0 {
		user.Course = body.Course
	}

	// 5. update the user record in the database
	result = initializers.DB.Model(&user).Updates(user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed updating the record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
