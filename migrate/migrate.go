package main

import (
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	// migrate the schemas
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Document{})
	initializers.DB.AutoMigrate(&models.Product{})
}
