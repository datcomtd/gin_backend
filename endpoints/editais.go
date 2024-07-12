package endpoints

import (
	"net/http"

	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"github.com/gin-gonic/gin"
)

type AddRequest struct {
	Body string

	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetEditais(c *gin.Context) {
	var editais []models.Document

	// encontra todos os editais
	initializers.DB.Find(&editais)

	c.JSON(http.StatusOK, gin.H{"document": editais})
}

func AddEdital(c *gin.Context) {
	var edital models.Document
	var body AddRequest

	// retrieve POST data
	c.Bind(&body)

	// title é um campo obrigatório
	if body.Title == "" {
		c.JSON(http.StatusBadGateway, gin.H{"message": "invalid request"})
		return
	}

	// verifica se o documento já existe
	result := initializers.DB.Where("title = ?", body.Title).First(&edital)
	if result.Error == nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "document already uploaded"})
		return
	}

	edital = models.Document{Title: body.Title, Description: body.Description}

	// salva no DB
	result = initializers.DB.Create(&edital)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create the record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"uploaded": edital})
}
