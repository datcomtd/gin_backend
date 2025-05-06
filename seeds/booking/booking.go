package booking

import (
    "time"
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"
)

var bookings = []models.Booking{
    {
        TimestampStart: time.Date(2025, 5, 6, 10, 0, 0, 0, time.UTC),
        TimestampEnd:   time.Date(2025, 5, 6, 12, 0, 0, 0, time.UTC),
        Description:    "Team Meeting",
        Day:           "Monday",
        Username:      "admin",
        Role:         1,
        Course:       1,
    },
    {
        TimestampStart: time.Date(2025, 5, 6, 14, 0, 0, 0, time.UTC),
        TimestampEnd:   time.Date(2025, 5, 6, 16, 0, 0, 0, time.UTC),
        Description:    "Project Discussion",
        Day:           "Monday",
        Username:      "john_doe",
        Role:         3,
        Course:       2,
    },
}

func LoadBookings() {
    for _, booking := range bookings {
        result := initializers.DB.Model(&models.Booking{}).Create(&booking)
        if result.Error != nil {
            panic(result.Error)
        }
    }
}