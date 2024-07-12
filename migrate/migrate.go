package main

import (
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	// migrate the schema
	initializers.DB.AutoMigrate(&models.Document{})
}
