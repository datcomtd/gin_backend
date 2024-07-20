package endpoints

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/authentication"
	"datcomtd/backend/initializers"
	"datcomtd/backend/models"
	"datcomtd/backend/utils"

	"net/http"
)

//
// GenerateKey(POST uploadRequest)
//  0. retrieve post data
//  1. check if the required fields are filled
//  2. check if the token is valid
//  3. check if the user has permission to upload a document
//  4. check if the file is already uploaded
//  5. create a new document record in the database
//  6. return the key to upload the file
//
// UploadDocument (POST f:file)
//  1. check if the token is valid
//  2. check if the user has permission to upload a document
//  3. check if the key is not empty
//  4. check if the key exists
//  5. check if the key is valid
//  6. get the file
//  7. save the file
//

type uploadRequest struct {
	Body string

	Title       string `json:"title"`
	Description string `json:"description"`

	Source   string `json:"source"`
	Category string `json:"category"`
}

func GenerateKey(c *gin.Context) {
	var body uploadRequest
	var document models.Document

	// 0. retrieve post data
	c.Bind(&body)

	// 1. check if the required fields are filled
	if (body.Title == "") || (body.Source == "") || (body.Category == "") {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required fields are not filled"})
		return
	}

	// 2. check if the token is valid and get the token's username
	username, userrole, errCode, errString := authentication.VerifyToken(c.GetHeader("Authorization"))
	if username == "" {
		c.JSON(errCode, gin.H{"message": errString})
		return
	}

	// 3. check if the user has permission to upload a document
	if userrole > 6 {
		c.JSON(http.StatusForbidden, gin.H{"message": "user does not have permission"})
		return
	}

	// 4. check if the file is already uploaded
	result := initializers.DB.Model(&models.Document{}).Where("title = ?", body.Title).First(&document)
	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "file already exists"})
		return
	}

	// 5. create a new document record in the database
	// 5.1. new document model
	document = models.Document{
		Key:           utils.RandomString(32),
		Title:         body.Title,
		Description:   body.Description,
		Source:        body.Source,
		Category:      body.Category,
		CreatedBy:     username,
		LastUpdatedBy: username,
	}
	// 5.2. insert the model into the database
	result = initializers.DB.Create(&document)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed creating the record"})
		return
	}

	// 6. return the key to upload the file
	c.JSON(http.StatusOK, gin.H{"key": document.Key})
}

func UploadDocument(c *gin.Context) {
	var document models.Document

	// 1. check if the token is valid
	username, userrole, errCode, errString := authentication.VerifyToken(c.GetHeader("Authorization"))
	if username == "" {
		c.JSON(errCode, gin.H{"message": errString})
		return
	}

	// 2. check if the user has permission to upload a document
	if userrole > 6 {
		c.JSON(http.StatusForbidden, gin.H{"message": "user does not have permission"})
		return
	}

	// 3. check if the key is not empty
	key := c.Param("key")
	if key == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid key"})
		return
	}

	// 4. check if the key exists
	result := initializers.DB.Model(&models.Document{}).Where("key = ?", c.Param("key")).First(&document)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid key"})
		return
	}

	// 5. check if the key is valid
	if key != document.Key {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid key"})
		return
	}

	// 6. get the file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid document"})
		return
	}

	// 7. save the file
	err = c.SaveUploadedFile(file, "./media/"+document.Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed saving the document"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"document": document})
}