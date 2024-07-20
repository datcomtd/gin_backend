package authentication

import (
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"net/http"
)

//
// VerifyToken()
//  1. check if the token is empty
//  2. check if the token exists, get user record
//  3. check if the token is valid
//  4. return the username and set the error to zero
//

func VerifyToken(token string) (string, int, string) {
	var user models.User

	// 1. check if the token is empty
	if token == "" {
		return "", http.StatusUnauthorized, "invalid token"
	}

	// 2. check if the token exists
	result := initializers.DB.Where("token = ?", token).First(&user)
	if result.Error != nil {
		return "", http.StatusUnauthorized, "invalid token"
	}

	// 3. check if the token is valid
	if token != user.Token {
		return "", http.StatusUnauthorized, "invalid token"
	}

	// 4. return the username and set the error to zero
	return user.Username, 0, ""
}
