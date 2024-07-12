package initializers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	DB, err = gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}
}
