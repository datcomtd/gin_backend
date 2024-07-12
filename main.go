package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"datcomtd/backend/endpoints"
	"datcomtd/backend/initializers"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	// CORS - Cross Origin Resource Share
	// default permite todas as origens
	router.Use(cors.Default())

	// Editais endpoints
	router.GET("/api/editais/", endpoints.GetEditais)
	router.POST("/api/editais/add", endpoints.AddEdital)

	router.Run("localhost:8000")
}
