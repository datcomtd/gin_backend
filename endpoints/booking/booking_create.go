package booking

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/authentication/token"
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"net/http"
	"time"
)

//
// Create (POST :createReq)
//  0. retrieve post data
//  1. check if the required fields are filled
//  2. check if the token is valid
//  3. check if the user has permission to create a booking
//  4. check if the booking already exists
//  5. create a new booking record in the database
//

type booking_createReq struct {
	Body string

	TimestampStart time.Time `json:"time-start"`
	TimestampEnd   time.Time `json:"time-end"`
	Description    string    `json:"description"`
}

func Create(c *gin.Context) {
	var body booking_createReq
	var booking models.Booking

	// 0. retrieve post data
	c.Bind(&body)

	// 1. check if the required fields are filled
	if body.TimestampStart.IsZero() || body.TimestampEnd.IsZero() || body.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required fields are not filled"})
		return
	}

	// 2. check if the token is valid
	username, userrole, usercourse, errCode, errString := token.VerifyToken(c.GetHeader("Authorization"))
	if username == "" {
		c.JSON(errCode, gin.H{"message": errString})
		return
	}

	// 3. check if the user has permission to create a booking
	if usercourse > initializers.ENUM_DATCOM_COURSE_MEMBER {
		c.JSON(http.StatusForbidden, gin.H{"message": "user does not have permission"})
		return
	}

	// 4. check if a booking is already taking place
	var count int64
	initializers.DB.Model(&models.Booking{}).Where("(start_time < ? AND end_time > ?) OR (start_time >= ? AND start_time < ?)",
		body.TimestampEnd, body.TimestampStart, body.TimestampStart, body.TimestampEnd).Count(&count)

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "a meeting is already booked"})
		return
	}

	// 5. create a new booking record in the database
	// 5.0. model
	booking = models.Booking{
		TimestampStart: body.TimestampStart,
		TimestampEnd:   body.TimestampEnd,

		Day: body.TimestampStart.Format("2001-01-01"),

		Description: body.Description,

		Username: username,
		Role:     userrole,
		Course:   usercourse,
	}
	// 5.1. create
	result := initializers.DB.Create(&booking)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed creating the record"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "booking created"})
}
