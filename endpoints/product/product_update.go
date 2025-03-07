package product

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/authentication/token"
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"net/http"
)

//
// UpdateProduct (POST :updateRequest)
//  0. retrieve post data
//  1. check if the token exists
//  2. check if the required fields are filled
//  3. check if the user has permission
//  4. check if the product exists
//  5. update the product fields
//  6. update the product record in the database
//

type product_updateRequest struct {
	Body string

	ID uint `json:"id"`

	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`

	Price   float64 `json:"price"`
	NoStock bool    `json:"no-stock"`
	Stock   bool    `json:"stock"`
}

func UpdateProduct(c *gin.Context) {
	var body product_updateRequest
	var product models.Product

	// 0. retrieve post data
	c.Bind(&body)

	// 1. check if the token exists
	username, userrole, errCode, errString := token.VerifyToken(c.GetHeader("Authorization"))
	if username == "" {
		c.JSON(errCode, gin.H{"message": errString})
		return
	}

	// 2. check if the required fields are filled
	if body.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required fields are not filled"})
		return
	}

	// 3. check if the user has permission
	if userrole > initializers.ENUM_DATCOM_ROLE_MEMBER {
		c.JSON(http.StatusForbidden, gin.H{"message": "user does not have permission"})
		return
	}

	// 4. check if the product exists
	result := initializers.DB.Model(&models.Product{}).Where("id = ?", body.ID).First(&product)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}

	// 5. update the product fields
	product.LastUpdatedBy = username
	// 5.1. title
	if body.Title != "" {
		product.Title = body.Title
	}
	// 5.2. description
	if body.Description != "" {
		product.Description = body.Description
	}
	// 5.3. category
	if body.Category != "" {
		product.Category = body.Category
	}
	// 5.4. price
	if body.Price != 0 {
		product.Price = body.Price
	}
	// 5.5. stock
	if body.NoStock != false {
		product.InStock = false
	}
	if body.Stock != false {
		product.InStock = true
	}

	// 6. update the product record in the database
	result = initializers.DB.Model(&product).Updates(product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed updating the record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}
