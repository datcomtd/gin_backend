package booking

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/authentication/token"
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"net/http"
)

//
// Delete (POST :deleteReq)
//  0. retrieve post data
//  1. check if the required fields are filled
//  2. check if the token is valid
//  3. get the booking record
//  4. check if the user has permission
//  5. delete the booking record
//

type booking_deleteReq struct {
	Body string

	ID uint `json:"id"`
}

func Delete(c *gin.Context) {
	var body booking_deleteReq
	var booking models.Booking

	// 0. retrieve post data
	c.Bind(&body)

	// 1. check if the required fields are filled
	if body.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required fields are not filled"})
		return
	}

	// 2. check if the token is valid
	username, _, _, errCode, errString := token.VerifyToken(c.GetHeader("Authorization"))
	if username == "" {
		c.JSON(errCode, gin.H{"message": errString})
		return
	}

	// 3. get the booking record
	result := initializers.DB.Model(&models.Booking{}).Where("id = ?", body.ID).First(&booking)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "booking not found"})
		return
	}

	// 4. check if the user has permission
	if booking.Username != username && username != initializers.Admin.Username {
		c.JSON(http.StatusForbidden, gin.H{"message": "user does not have permission"})
		return
	}

	// 5. delete the booking record
	result = initializers.DB.Delete(&booking)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed deleting the record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "booking deleted"})
}

