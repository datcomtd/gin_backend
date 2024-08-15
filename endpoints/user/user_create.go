package user

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
//  2. check if the admin username is valid
//  3. check if the admin password is correct
//  4. check if the course value is valid
//  5. check if the user is already registered
//  6. hash the body password
//  7. create a record in the database
//  8. if role is President, update ADMIN credentials
//

type user_registerRequest struct {
	Body string

	AdminUsername string `json:"admin-username"`
	AdminPassword string `json:"admin-password"`

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
	if (body.Password == "") || (body.Username == "") || (body.Role == 0) || (body.Course == 0) ||
		(body.AdminUsername == "") || (body.AdminPassword == "") {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required fields are not filled"})
		return
	}

	// 1. check if the admin username is valid
	if body.AdminUsername != initializers.Admin.Username {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid admin username or password"})
		return
	}

	// 2. check if the admin password is correct
	// 2.0. hash the body password
	bl := authentication.VerifyPassword(body.AdminPassword, initializers.Admin.Password)
	if bl != true {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid admin username or password"})
		return
	}

	// 4. check if the course value is valid
	if body.Course > 2 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid course"})
		return
	}

	// 5. check if the user is already registered
	result := initializers.DB.Where("username = ?", strings.ToLower(body.Username)).First(&user)
	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user is already registered"})
		return
	}

	// 6. hash the body password
	hashedPassword, err := authentication.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed hashing the password"})
		return
	}

	// 7. create the record in the database
	// 7.1. create the model for insertion
	user = models.User{
		Token_UpdatedAt: time.Now().UTC(),
		Token:           utils.RandomString(64),
		Password:        hashedPassword,

		Username: strings.ToLower(body.Username),
		Email:    body.Email,

		Role:   body.Role,
		Course: body.Course,
	}
	// 7.2. insert the model into the database
	result = initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed creating the record"})
		return
	}

	// 8. if role is President, update ADMIN credentials
	if user.Role == models.President {
		initializers.Admin = &user
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}
