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
// DeleteUser (POST deleteRequest)
//  0. retrieve post data
//  1. check if the required fields are filled
//  2. check if the user exists
//  3. check if the password is correct
//  4. delete the user record
//

type user_deleteRequest struct {
	Body string

	Username string `json:"username"`
	Password string `json:"password"`
}

func DeleteUser(c *gin.Context) {
	var body user_deleteRequest
	var user models.User

	// 0. retrieve post data
	c.Bind(&body)

	// 1. check if the required fields are filled
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

	// 4. delete the user record
	result = initializers.DB.Delete(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed deleting the record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
