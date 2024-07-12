package endpoints

import (
	"net/http"

	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"github.com/gin-gonic/gin"
)

func GetEditais(c *gin.Context) {
	var editais []models.Document

	// encontra todos os editais
	initializers.DB.Find(&editais)

	c.JSON(http.StatusOK, gin.H{"document": editais})
}
