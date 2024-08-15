package main

import (
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"
)

func init() {
	initializers.ConnectToDB()
	initializers.SetupAdmin()
}

func main() {
	// migrate the schemas
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Document{})
	initializers.DB.AutoMigrate(&models.Product{})
}
