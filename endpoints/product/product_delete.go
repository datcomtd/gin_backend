package product

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/authentication/token"
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"net/http"
)

//
// DeleteProduct (POST :deleteRequest)
//  0. retrieve post data
//  1. check if the required fields are filled
//  2. check if the token is valid
//  3. check if the user has permission
//  4. get the product record
//  5. delete the product record
//

type product_deleteRequest struct {
	Body string

	ID uint `json:"id"`
}

func DeleteProduct(c *gin.Context) {
	var body product_deleteRequest
	var product models.Product

	// 0. retrieve post data
	c.Bind(&body)

	// 1. check if the required fields are filled
	if body.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required fields are not filled"})
		return
	}

	// 2. check if the token is valid
	username, userrole, errCode, errString := token.VerifyToken(c.GetHeader("Authorization"))
	if username == "" {
		c.JSON(errCode, gin.H{"message": errString})
		return
	}

	// 3. check if the user has permission
	if userrole > initializers.ENUM_DATCOM_ROLE_MEMBER {
		c.JSON(http.StatusForbidden, gin.H{"message": "user does not have permission"})
		return
	}

	// 4. get the product record
	result := initializers.DB.Model(&models.Product{}).Where("id = ?", body.ID).First(&product)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}

	// 5. delete the product record
	result = initializers.DB.Delete(&product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed deleting the record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}
