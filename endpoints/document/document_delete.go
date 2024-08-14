package document

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/authentication"
	"datcomtd/backend/authentication/token"
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
//  4. check if the token's user is the document creator or ADMIN authority
//  5. delete the document file
//  6. delete the document record
//

type document_deleteRequest struct {
	Body string

	AdminUsername string `json:"admin-username"`
	AdminPassword string `json:"admin-password"`

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
	username, _, errCode, errString := token.VerifyToken(c.GetHeader("Authorization"))
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

	// 4. check if the token's user is the document creator or ADMIN authority
	if body.AdminPassword == "" {
		if username != document.CreatedBy {
			c.JSON(http.StatusForbidden, gin.H{"message": "user is not the document's creator"})
			return
		}
	} else {
		// 4.B. ADMIN authority
		if body.AdminUsername != initializers.Admin.Username {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid admin username or password"})
			return
		}
		bl := authentication.VerifyPassword(body.AdminPassword, initializers.Admin.Password)
		if bl != true {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid admin username or password"})
			return
		}
	}

	// 5. delete the document file
	err := os.Remove("./media/document/" + document.Key + "_" + document.Filename)
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
