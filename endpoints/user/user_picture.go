package user

import (
	"github.com/gin-gonic/gin"

	"datcomtd/backend/authentication/token"

	"net/http"
	"path/filepath"
)

//
// UploadPicture (POST _FormFile:file)
//  1. check if the token is valid
//  2. get the file
//  3. check if the file is valid
//  4. save the file
//

func UploadPicture(c *gin.Context) {

	// 1. check if the token is valid
	username, _, errCode, errString := token.VerifyToken(c.GetHeader("Authorization"))
	if username == "" {
		c.JSON(errCode, gin.H{"message": errString})
		return
	}

	// 2. get the file
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid file"})
		return
	}

	extension := filepath.Ext(file.Filename)
	save_path := "./media/member/" + username + extension

	// 3. check if the file is valid
	if extension != ".png" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid extension"})
		return
	}

	// 4. save the file
	err = c.SaveUploadedFile(file, save_path)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed uploading the file"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "picture uploaded"})
}
