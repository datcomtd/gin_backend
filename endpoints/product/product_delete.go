package product

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

func DeleteProduct(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented yet"})
}
