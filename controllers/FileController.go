package controllers

import (
	initializers "Basic/Auth-Api/Initializers"
	models "Basic/Auth-Api/Models"
	token "Basic/Auth-Api/Token"
	"io"
	"net/http"
	"net/url"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

var (
	storageClient *storage.Client
)

// HandleFileUploadToBucket uploads file to bucket
func UpdateProfileImg(c *gin.Context) {
	bucket := "go-first-user.appspot.com" //your bucket name

	user_id, errT := token.ExtractTokenID(c)

	if errT != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errT.Error()})
		return
	}

	var err error

	ctx := appengine.NewContext(c.Request)

	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	f, fileUploaded, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	defer f.Close()

	sw := storageClient.Bucket(bucket).Object(fileUploaded.Filename).NewWriter(ctx)

	if _, err := io.Copy(sw, f); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	if err := sw.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	u, err := url.Parse("/" + bucket + "/" + sw.Attrs().Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"Error":   true,
		})
		return
	}
	initializers.DB.Model(&models.User{}).Where("id = ?", user_id).Update("profile_img", u.EscapedPath())

	c.JSON(http.StatusOK, gin.H{
		"message": "profile image successfully updated",
	})
}

func RemoveProfileImg(c *gin.Context) {
	user_id, errT := token.ExtractTokenID(c)

	if errT != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errT.Error()})
		return
	}
	initializers.DB.Model(&models.User{}).Where("id = ?", user_id).Update("profile_img", "")

	c.JSON(http.StatusOK, gin.H{
		"message": "profile image successfully Deleted",
	})
}
