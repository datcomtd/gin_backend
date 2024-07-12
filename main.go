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

	// Editais endpoints
	router.GET("/api/editais/", endpoints.GetEditais)
	router.POST("/api/editais/add", endpoints.AddEdital)

	router.Run("localhost:8000")
}
