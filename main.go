package main

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/endpoints"
	"datcomtd/backend/initializers"
	"datcomtd/backend/utils"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	// CORS - Cross Origin Resource Share
	// default permite todas as origens
	router.Use(utils.CORSMiddleware())

	// Authentication endpoints
	router.POST("/api/register", endpoints.Register)
	router.POST("/api/token", endpoints.GetToken)

	// Members endpoints
	router.GET("/api/users", endpoints.GetUsers)
	router.GET("/api/user/:username", endpoints.GetUserByUsername)

	router.Run("localhost:8000")
}
