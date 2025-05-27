package product

import (
	"slices"

	"github.com/gin-gonic/gin"

	"datcomtd/backend/authentication/token"
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"net/http"
	"os"
	"path/filepath"
)

//
// PhotoAdd (PARAM :id/:name, _FileForm)
//  0. check if the parameters are filled
//  1. check if the token is valid
//  2. check if the user has permission to upload a product photo
//  3. check if the product exists
//  4. get the file
//  5. check if the file is valid
//  6. save the file
//  7. update the product record
//
// PhotoDelete (PARAM :id/:name)
//  0. check if the parameters are filled
//  1. check if the token is valid
//  2. check if the user has permission to delete a product photo
//  3. check if the product exists
//  4. delete the photo file
//  5. delete the photo_name from the product.Photos
//

func PhotoAdd(c *gin.Context) {
	var product models.Product

	product_id := c.Param("id")
	photo_name := c.Param("name")

	// 0. check if the parameters are filled
	if product_id == "" || photo_name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required parameters are not filled"})
		return
	}

	// 1. check if the token is valid
	username, userrole, usercourse, errCode, errString := token.VerifyToken(c.GetHeader("Authorization"))
	if username == "" {
		c.JSON(errCode, gin.H{"message": errString})
		return
	}

	// 2. check if the user has permission to upload a product photo
	if userrole > initializers.ENUM_DATCOM_ROLE_MEMBER || usercourse > initializers.ENUM_DATCOM_COURSE_MEMBER {
		c.JSON(http.StatusForbidden, gin.H{"message": "user does not have permission"})
		return
	}

	// 3. check if the product exists
	result := initializers.DB.Model(&models.Product{}).Where("id = ?", product_id).First(&product)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}

	// 4. get the file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	extension := filepath.Ext(file.Filename)
	save_path := "./media/product/" + product_id + "-" + photo_name + extension

	// 5. check if the file is valid (png, jpg, jpeg)
	if extension != ".png" && extension != ".jpg" && extension != ".jpeg" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid extension"})
		return
	}

	// 6. save the file
	err = c.SaveUploadedFile(file, save_path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed uploading the file"})
		return
	}

	// 7. update the product record
	product.Count = product.Count + 1
	product.Photos = append(product.Photos, save_path)
	// 7.B. update
	result = initializers.DB.Model(&product).Updates(product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed updating the record"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "photo uploaded"})
}

func PhotoDelete(c *gin.Context) {
	var product models.Product

	product_id := c.Param("id")
	photo_name := c.Param("name")

	// 0. check if the parameters are filled
	if product_id == "" || photo_name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required parameters are not filled"})
		return
	}

	// 1. check if the token is valid
	username, userrole, usercourse, errCode, errString := token.VerifyToken(c.GetHeader("Authorization"))
	if username == "" {
		c.JSON(errCode, gin.H{"message": errString})
		return
	}

	// 2. check if the user has permission to delete a product photo
	if userrole > initializers.ENUM_DATCOM_ROLE_MEMBER || usercourse > initializers.ENUM_DATCOM_COURSE_MEMBER {
		c.JSON(http.StatusForbidden, gin.H{"message": "user does not have permission"})
		return
	}

	// 3. check if the product exists
	result := initializers.DB.Model(&models.Product{}).Where("id = ?", product_id).First(&product)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}

	// 4. delete the photo file
	save_path := "./media/product/" + product_id + "-" + photo_name
	err := os.Remove(save_path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// 5. delete the photo_name from the product.Photos
	var i int
	for index, photo := range product.Photos {
		if photo == save_path {
			i = index
		}
	}
	product.Count = product.Count - 1
	product.Photos = slices.Delete(product.Photos, i, len(product.Photos))
	// 5.B. update
	result = initializers.DB.Model(&product).Updates(product)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed updating the record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "photo deleted"})
}
