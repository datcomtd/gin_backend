package token

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/authentication"
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"
	"datcomtd/backend/utils"

	"net/http"
	"time"
)

//
// GetToken (POST TokenRequest)
//  0. retrieve post data
//  1. check if the required fields are filled
//  2. check if the user exists
//  3. check if the password is correct
//  4. check if the user's token is expired, if so, generate a new one
//

type TokenRequest struct {
	Body string

	Username string `json:"username"`
	Password string `json:"password"`
}

func GetToken(c *gin.Context) {
	var body TokenRequest
	var user models.User

	// 0. retrieve post data
	c.Bind(&body)

	// 1. check if the required fields are filled
	if (body.Username == "") || (body.Password == "") {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required fields are not filled"})
		return
	}

	// 2. check if the user exists
	// 2.1. sql query
	result := initializers.DB.Model(&models.User{}).Where("username = ?", body.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid username or password"})
		return
	}
	// 2.2. check if the username is valid
	if body.Username != user.Username {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid username or password"})
		return
	}

	// 3. check if the password is correct
	bl := authentication.VerifyPassword(body.Password, user.Password)
	if bl != true {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid username or password"})
		return
	}

	// 4. check if the user's token is expired, if so, generate a new one
	if time.Now().UTC().Sub(user.Token_UpdatedAt).Hours() > initializers.TOKEN_EXPIRES_HOURS {
		// 4-A.1. generate a new token
		user.Token_UpdatedAt = time.Now().UTC()
		user.Token = utils.RandomString(64)
		// 4-A.2. update the user record in the database
		result = initializers.DB.Model(&user).Updates(user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed updating the record"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"token": user.Token})
}
