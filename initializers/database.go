package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"datcomtd/backend/authentication"
	"datcomtd/backend/models"
	"datcomtd/backend/utils"

	"time"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	DSN := "host=localhost user=datcom_user password=" + DATCOM_DB_PWD + " dbname=datcom_db port=4145 sslmode=disable TimeZone=Europe/London"

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  DSN,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
}

var Admin *models.User

func SetupAdmin() {
	result := DB.Model(&models.User{}).Where("role = ?", models.President).First(&Admin)
	if result.Error != nil {
		hashedPassword, err := authentication.HashPassword(DATCOM_ADMIN_PWD)
		if err != nil {
			panic("failed hashing the admin password")
		}

		Admin.Token_UpdatedAt = time.Now().UTC()
		Admin.Token = utils.RandomString(64)

		Admin.Username = DATCOM_ADMIN_USER
		Admin.Password = hashedPassword
		Admin.Role = models.President
	}
}
