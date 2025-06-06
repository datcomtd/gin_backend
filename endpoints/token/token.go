package token

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/authentication"
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"
	"github.com/golang-jwt/jwt/v5"

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

	Email string `json:"email"`
	Password string `json:"password"`
}

func GetToken(c *gin.Context) {
	var body TokenRequest
	var user models.User

	// 0. retrieve post data
	c.Bind(&body)

	// 1. check if the required fields are filled
	if (body.Email == "") || (body.Password == "") {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required fields are not filled"})
		return
	}

	// 2. check if the user exists
	// 2.1. sql query
	result := initializers.DB.Model(&models.User{}).Where("email = ?", body.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid email or password"})
		return
	}
	// 2.2. check if the email is valid
	if body.Email != user.Email {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid email or password"})
		return
	}

	// 3. check if the password is correct
	bl := authentication.VerifyPassword(body.Password, user.Password)
	if bl != true {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid email or password"})
		return
	}

	// generate a new token jwt
	token, err := GenerateJWT(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}


func GenerateJWT(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email": email,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour).Add(time.Hour).Add(time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("Datcom@td2025#"))
}