package users

import (
    "datcomtd/backend/initializers"
	"datcomtd/backend/models"
    "golang.org/x/crypto/bcrypt"
    "log"
)

// generate hashed passwords
func hashPassword(password string) string {
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        log.Fatalf("failed to hash password: %v", err)
    }
    return string(hashed)
}

var users = []models.User{
    {
        Username: "patrick",
        Email:    "patrick@example.com",
        Role:     1,
        Course:   1,
        RA:       "RA123",
        Password: hashPassword("patrick123"), // Should be hashed in production
    },
    {
        Username: "lucas",
        Email:    "lucas@example.com",
        Role:     2,
        Course:   2,
        RA:       "RA456",
        Password: hashPassword("lucas123"), // Should be hashed in production
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