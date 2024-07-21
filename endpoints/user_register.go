package endpoints

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/authentication"
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"
	"datcomtd/backend/utils"

	"net/http"
	"strings"
	"time"
)

//
// Register (POST registerRequest)
//  0. retrieve post data
//  1. check if the required fields are filled
//  2. check if the course value is valid
//  3. check if the user is already registered
//  4. hash the body password
//  5. create a record in the database
//

type user_registerRequest struct {
	Body string

	Password string `json:"password"`

	Username string `json:"username"`
	Email    string `json:"email"`

	Role   uint `json:"role"`
	Course uint `json:"course"`
}

func Register(c *gin.Context) {
	var body user_registerRequest
	var user models.User

	// 0. retrieve post data
	c.Bind(&body)

	// 1. check if the required fields are filled
	if (body.Password == "") || (body.Username == "") || (body.Role == 0) || (body.Course == 0) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required fields are not filled"})
		return
	}

	// 2. check if the course value is valid
	if body.Course > 2 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid course"})
		return
	}

	// 3. check if the user is already registered
	result := initializers.DB.Where("username = ?", strings.ToLower(body.Username)).First(&user)
	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user is already registered"})
		return
	}

	// 4. hash the body password
	hashedPassword, err := authentication.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed hashing the password"})
		return
	}

	// 5. create the record in the database
	// 5.1. create the model for insertion
	user = models.User{
		Token_UpdatedAt: time.Now().UTC(),
		Token:           utils.RandomString(64),
		Password:        hashedPassword,

		Username: strings.ToLower(body.Username),
		Email:    body.Email,

		Role:   body.Role,
		Course: body.Course,
	}
	// 5.2. insert the model into the database
	result = initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed creating the record"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}
