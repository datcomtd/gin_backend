package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
