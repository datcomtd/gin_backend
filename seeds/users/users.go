package users

import (
    "datcomtd/backend/initializers"
	"datcomtd/backend/models"
)

var users = []models.User{
    {
        Username: "admin",
        Email:    "admin@example.com",
        Role:     1,
        Course:   1,
        RA:       "RA123",
        Password: "admin123", // Should be hashed in production
    },
    {
        Username: "john_doe",
        Email:    "john@example.com",
        Role:     3,
        Course:   2,
        RA:       "RA456",
        Password: "test123", // Should be hashed in production
    },
}

func LoadUsers() {
    for _, user := range users {
        result := initializers.DB.Model(&models.User{}).Create(&user)
        if result.Error != nil {
            panic(result.Error)
        }
    }
}