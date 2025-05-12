package token

import (
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"
	"github.com/golang-jwt/jwt/v5"

	"errors"
	"net/http"
)

//
// VerifyToken()
//  1. check if the token is empty
//  2. check if the token exists, get user record
//  3. check if the token is valid
//  4. return the username and set the error to zero
//

func VerifyToken(token string) (string, uint, uint, int, string) {
	var user models.User

	// 1. check if the token is empty
	if token == "" {
		return "", 0, 0, http.StatusUnauthorized, "invalid token"
	}

	// 2. check if the token is valid 
	userId, status, err := ValidateJWT(token)
	if userId == 0 {
		return "", 0, 0, status, err
	}

	// 2. check if the token exists
	result := initializers.DB.Where("id = ?", userId).First(&user)
	if result.Error != nil {
		return "", 0, 0, http.StatusUnauthorized, "invalid token"
	}

	// 4. return the username and set the error to zero
	return user.Username, user.Role, user.Course, 0, ""
}


func ValidateJWT(tokenString string) (float64, int, string) {
	if tokenString == "" {
		return 0, 401, "required token"
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		// Return the secret key
		return []byte("Datcom@td2025#"), nil
	})


	// Check for parsing errors
	if err != nil {
		return 0, 401, err.Error()
	}

	// Check if the token is valid and extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Extract the user ID
		userId := claims["user_id"].(float64)
		return userId, 200, ""
	}
	return 0, 401, "invalid token3"
}