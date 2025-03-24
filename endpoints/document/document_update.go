package document

import (
	"datcomtd/backend/authentication/token"
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//
// UpdateDocument (POST document_updateRequest)
//  0. retrieve post data
//  1. check if the token exists
//  2. check if the required fields are filled
//  3. check if the user has permission
//  4. check if the document exists
//  5. update the document's fields
//  6. update the document record in the database
//

type document_updateRequest struct {
	Body string

	ID uint `json:"id"`

	Filename    string `json:"filename"`
	Title       string `json:"title"`
	Description string `json:"description"`

	Source   string `json:"source"`
	Category string `json:"category"`
}

func UpdateDocument(c *gin.Context) {
	var body document_updateRequest
	var document models.Document

	// 0. retrieve post data
	c.Bind(&body)

	// 1. check if the token exists
	username, userrole, usercourse, errCode, errString := token.VerifyToken(c.GetHeader("Authorization"))
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
	if userrole > initializers.ENUM_DATCOM_ROLE_MEMBER || usercourse > initializers.ENUM_DATCOM_COURSE_MEMBER {
		c.JSON(http.StatusForbidden, gin.H{"message": "user does not have permission"})
		return
	}

	// 4. check if the document exists
	result := initializers.DB.Model(&models.Document{}).Where("id = ?", body.ID).First(&document)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "document not found"})
		return
	}

	// 5. update the document's fields
	// 5.1. filename
	if (body.Filename != "") && (body.Filename != document.Filename) {
		// 5.1.1. rename the file
		err := os.Rename("./media/"+document.Key+"_"+document.Filename,
			"./media/"+document.Key+"_"+body.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed renaming the file"})
			return
		}
		// 5.1.2. change the document filename field
		document.Filename = body.Filename
	}
	// 5.2. title
	if (body.Title != "") && (body.Title != document.Title) {
		document.Title = body.Title
	}
	// 5.3. description
	if (body.Description != "") && (body.Description != document.Description) {
		document.Description = body.Description
	}
	// 5.4. source
	if (body.Source != "") && (body.Source != document.Source) {
		document.Source = body.Source
	}
	// 5.5. category
	if (body.Category != "") && (body.Category != document.Category) {
		document.Category = body.Category
	}

	// 6. update the document record in the database
	// 6.0. change the document LastUpdatedBy field first
	document.LastUpdatedBy = username
	result = initializers.DB.Model(&document).Updates(document)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed updating the record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"document": document})
}
