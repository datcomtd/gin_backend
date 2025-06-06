package product

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/authentication/token"
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"
	"time"
	"fmt"

	"net/http"
)

//
// CreateProduct (POST :createRequest)
//  0. retrieve post data
//  1. check if the required fields are filled
//  2. check if the token is valid
//  3. check if the user has permission to create a product
//  4. check if the product already exists
//  5. create a new product record in the database
//

type product_createRequest struct {
	Body string

	Title       string 	`json:"title"`
	Description string 	`json:"description"`
	Category    string 	`json:"category"`
	Count	    int		`json:"count"`
	InStock     bool    `json:"in_stock"`


	Price float64 `json:"price"`
}

func CreateProduct(c *gin.Context) {
	var body product_createRequest
	var product models.Product

	// 0. retrieve post data
	c.Bind(&body)

	// 1. check if the required fields are filled
	if body.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required fields are not filled"})
		return
	}

	// 2. check if the token is valid
	username, userrole, usercourse, errCode, errString := token.VerifyToken(c.GetHeader("Authorization"))
	if username == "" {
		c.JSON(errCode, gin.H{"message": errString})
		return
	}

	// 3. check if the user has permission to create a product
	if userrole > initializers.ENUM_DATCOM_ROLE_MEMBER || usercourse > initializers.ENUM_DATCOM_COURSE_MEMBER {
		c.JSON(http.StatusForbidden, gin.H{"message": "user does not have permission"})
		return
	}

	// 4. check if the product already exists
	result := initializers.DB.Model(&models.Product{}).Where("title = ?", body.Title).First(&product)
	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "product already exists"})
		return
	}

	// print body
	fmt.Println(body)

	// 5. create a new product record in the database
	// 5.0. model
	product = models.Product{
		Title:       body.Title,
		Description: body.Description,
		Category:    body.Category,

		Price:   	 body.Price,
		InStock:     body.InStock,
		Count:   	 body.Count,

		CreatedBy:     username,
		LastUpdatedBy: username,
		UpdatedAt:    time.Now(),
		CreatedAt:   time.Now(),
	}
	// 5.1. create
	result = initializers.DB.Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed creating the record"})
		return
	}

	c.JSON(200, gin.H{"message": "product created", "id": product.ID})
}
