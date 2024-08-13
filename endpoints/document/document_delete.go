package document

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/authentication"
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"

	"net/http"
	"os"
)

//
// DeleteDocument (POST document_deleteRequest)
//  0. retrieve post data
//  1. check if the required fields are filled
//  2. check if the token is valid
//  3. get the document record
//  4. check if the token's user is the document creator or ADMIN-KEY
//  5. delete the document file
//  6. delete the document record
//

type document_deleteRequest struct {
	Body string

	AdminKey string `json:"admin-key"`

	ID uint `json:"id"`
}

func DeleteDocument(c *gin.Context) {
	var body document_deleteRequest
	var document models.Document

	// 0. retrieve post data
	c.Bind(&body)

	// 1. check if the required fields are filled
	if body.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required fields are not filled"})
		return
	}

	// 2. check if the token is valid
	username, _, errCode, errString := authentication.VerifyToken(c.GetHeader("Authorization"))
	if username == "" {
		c.JSON(errCode, gin.H{"message": errString})
		return
	}

	// 3. get the document record
	result := initializers.DB.Model(&models.Document{}).Where("id = ?", body.ID).First(&document)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "document not found"})
		return
	}

	// 4. check if the token's user is the document creator or ADMIN-KEY
	if username != document.CreatedBy && body.AdminKey != initializers.DATCOM_ADMIN_KEY {
		c.JSON(http.StatusForbidden, gin.H{"message": "user is not the document's creator"})
		return
	}

	// 5. delete the document file
	err := os.Remove("./media/" + document.Key + "_" + document.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed deleting the document"})
		return
	}

	// 6. delete the document record
	result = initializers.DB.Delete(&document)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed deleting the record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "document deleted"})
}
