package user

import (
	"github.com/gin-gonic/gin"
	"datcomtd/backend/authentication/token"
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
//  3. check if the password is correct or ADMIN authority
//  4. update the user struct fields
//  5. update the user record in the database
//

type user_updateRequest struct {
	Body string

	AdminUsername string `json:"admin-username"`
	AdminPassword string `json:"admin-password"`

	Username string `json:"username"`
	NewUsername string `json:"newusername"`
	Password string `json:"password"`

	NewPassword string `json:"newpassword"`
	Email       string `json:"email"`

	Role   uint   `json:"role"`
	Course uint   `json:"course"`
	RA     string `json:"ra"`
}

func UpdateUser(c *gin.Context) {
    var body user_updateRequest
    var user models.User

    // 0. Retrieve post data
    if err := c.Bind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
        return
    }

    // 1. Check if the username is provided
    if body.Username == "" {
        c.JSON(http.StatusBadRequest, gin.H{"message": "username is required"})
        return
    }

    // 2. Verify the token
    requesterUsername, requesterRole, requesterCourse, errCode, errString := token.VerifyToken(c.GetHeader("Authorization"))
    if requesterUsername == "" {
        c.JSON(errCode, gin.H{"message": errString})
        return
    }

    // 3. Check if the user exists
    result := initializers.DB.Where("username = ?", strings.ToLower(body.Username)).First(&user)
    if result.Error != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"message": "user not found"})
        return
    }

    // 4. Permission checks
    if requesterCourse != 1 && requesterCourse != 2 {
        // Not in course 1 or 2: can only update own data
        if requesterUsername != user.Username {
            c.JSON(http.StatusForbidden, gin.H{"message": "you can only update your own data"})
            return
        }
    } else {
        // In course 1 or 2
        if requesterRole > 1 {
            // Role > 1: can only update own data
            if requesterUsername != user.Username {
                c.JSON(http.StatusForbidden, gin.H{"message": "you can only update your own data"})
                return
            }
        }
        // Role == 1: can update any user's data (no restriction needed here)
    }

    // 5. Update the user fields
    // 5.1. New password
    if body.NewPassword != "" {
        hashedPassword, err := authentication.HashPassword(body.NewPassword)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to hash password"})
            return
        }
        user.Password = hashedPassword
    }
    // 5.2. Email
    if body.Email != "" {
        user.Email = body.Email
    }
    // 5.3. RA
    if body.RA != "" {
        user.RA = body.RA
    }
    // 5.4. Role and Course (only if requester is in course 1 or 2 with role 1)
    if (requesterCourse == 1 || requesterCourse == 2) && requesterRole == 1 {
        if body.Role != 0 {
            user.Role = body.Role
        }
        if body.Course != 0 {
            user.Course = body.Course
        }
    }
	if body.NewUsername != "" {
		// replace blanck spaces with underscores and trim the string
		body.NewUsername = strings.TrimSpace(body.NewUsername)
		body.NewUsername = strings.ReplaceAll(body.NewUsername, " ", "-")
		user.Username = strings.ToLower(body.NewUsername)
	}

    // 6. Update the user record in the database
    result = initializers.DB.Save(&user)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to update user"})
        return
    }

    // 7. Return the updated user
    c.JSON(http.StatusOK, gin.H{"user": user})
}
