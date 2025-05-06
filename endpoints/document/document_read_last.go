package document

import (
	"github.com/gin-gonic/gin"
	"datcomtd/backend/authentication/token"
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"net/http"
)

// GetLastTenDocuments returns the ten most recently created documents
func GetLastTenDocuments(c *gin.Context) {
	var documents []models.Document

	username, _, _, errCode, errString := token.VerifyToken(c.GetHeader("Authorization"))
	if username == "" {
		c.JSON(errCode, gin.H{"message": errString})
		return
	}

	result := initializers.DB.Model(&models.Document{}).Order("created_at DESC").Limit(10).Find(&documents)

	c.JSON(http.StatusOK, gin.H{
		"count":    result.RowsAffected,
		"document": documents})
	

}