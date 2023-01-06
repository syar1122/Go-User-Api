package controllers

import (
	initializers "Basic/Auth-Api/Initializers"
	models "Basic/Auth-Api/Models"
	"log"

	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	var users []models.User

	initializers.DB.Find(&users)

	c.JSON(200, &users)
}

func PostUser(c *gin.Context) {

	var body struct {
		FullName   string
		Password   string
		Username   string
		ProfileImg string
		RoleId     int
		Status     models.Status
	}
	c.Bind(&body)
	user := models.User{FullName: body.FullName, Password: body.Password, Username: body.Username, ProfileImg: body.ProfileImg, RoleId: body.RoleId, Status: int8(body.Status)}
	result := initializers.DB.Create(&user).Preload("Role").First(&user)

	if result.Error != nil {
		c.Status(400)
		log.Panic(result.Error)
		return
	}

	c.JSON(200, gin.H{
		"user": &user,
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	initializers.DB.Preload("Role").First(&user, id)

	c.JSON(200, gin.H{
		"user": &user,
	})
}

func UpdateUser(c *gin.Context) {
	var body struct {
		FullName   string
		Password   string
		Username   string
		ProfileImg string
		RoleId     int
		Status     int
	}

	id := c.Param("id")
	c.Bind(&body)
	var user models.User
	// userUpdate := models.User{FullName: body.FullName, Password: body.Password, Username: body.Username, ProfileImg: body.ProfileImg, RoleId: body.RoleId, Status: int8(body.Status)}
	initializers.DB.First(&user, id)
	result := initializers.DB.Model(&user).Updates(models.User{FullName: body.FullName, Password: body.Password, Username: body.Username, ProfileImg: body.ProfileImg, RoleId: body.RoleId, Status: int8(body.Status)})

	if result.Error != nil {
		c.Status(400)
		panic(result.Error)
	}

	c.JSON(200, gin.H{
		"user": &user,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.User{}, id)
	c.Status(200)
}
