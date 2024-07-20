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

	// User endpoints
	router.GET("/api/users", endpoints.GetUsers)
	router.GET("/api/user/:username", endpoints.GetUserByUsername)
	router.POST("/api/user/update", endpoints.UpdateUser)
	router.POST("/api/user/delete", endpoints.DeleteUser)

	// Document endpoints
	router.GET("/api/documents", endpoints.GetDocuments)
	router.GET("/api/document/by-id/:id", endpoints.GetDocumentByID)
	router.GET("/api/document/by-category/:category", endpoints.GetDocumentsByCategory)
	router.POST("/api/document/upload", endpoints.GenerateKey)
	router.POST("/api/document/upload/:key", endpoints.UploadDocument)
	router.POST("/api/document/delete", endpoints.DeleteDocument)

	router.Run("localhost:8000")
}
