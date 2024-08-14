package product

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func UpdateProduct(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented yet"})
}
