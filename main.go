package main

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/endpoints/booking"
	"datcomtd/backend/endpoints/document"
	"datcomtd/backend/endpoints/product"
	"datcomtd/backend/endpoints/token"
	"datcomtd/backend/endpoints/user"
	"datcomtd/backend/endpoints/dashboard"

	"datcomtd/backend/initializers"
	"datcomtd/backend/utils"
)

func init() {
	initializers.ConnectToDB()
	initializers.SetupAdmin()
}

func main() {
	router := gin.Default()

	router.Static("/media", "./media")

	// CORS - Cross Origin Resource Share
	// default permite todas as origens
	router.Use(utils.CORSMiddleware())

	// Authentication endpoints
	router.POST("/api/register", user.Register)
	router.POST("/api/token", token.GetToken)

	// User endpoints
	router.GET("/api/users", user.GetUsers)
	router.GET("/api/user/:username", user.GetUserByUsername)
	router.POST("/api/user/update", user.UpdateUser)
	router.POST("/api/user/delete", user.DeleteUser)
	router.POST("/api/user/picture", user.UploadPicture)
	router.POST("/api/user/create", user.Create)

	// Document endpoints
	router.GET("/api/documents", document.GetDocuments)
	router.GET("/api/document/by-id/:id", document.GetDocumentByID)
	router.GET("/api/document/by-category/:category", document.GetDocumentsByCategory)
	router.POST("/api/document/upload", document.GenerateKey)
	router.POST("/api/document/upload/:key", document.UploadDocument)
	router.POST("/api/document/update", document.UpdateDocument)
	router.POST("/api/document/delete", document.DeleteDocument)

	// Product endpoints
	router.GET("/api/products", product.GetProducts)
	router.GET("/api/product/by-id/:id", product.GetProductByID)
	router.GET("/api/product/by-category/:category", product.GetProductByCategory)
	router.POST("/api/product/create", product.CreateProduct)
	router.POST("/api/product/update", product.UpdateProduct)
	router.POST("/api/product/delete", product.DeleteProduct)
	router.POST("/api/product/:id/photo/:name/add", product.PhotoAdd)
	router.GET("/api/product/:id/photo/:name/delete", product.PhotoDelete)

	// Booking endpoints
	router.GET("/api/bookings", booking.View)
	router.GET("/api/booking/by-day/:day", booking.ViewByDay)
	router.GET("/api/booking/by-id/:id", booking.ViewById)
	router.POST("/api/booking/create", booking.Create)
	router.POST("/api/booking/delete", booking.Delete)

	// dashboard endpoints
	router.GET("/api/booking/week", booking.GetCurrentWeekBookings)
	router.GET("/api/document/last", document.GetLastTenDocuments)
	router.GET("/api/dashboard/counters", dashboard.GetCounts)

	router.Run("0.0.0.0:8000")
}
