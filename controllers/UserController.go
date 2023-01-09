package controllers

import (
	initializers "Basic/Auth-Api/Initializers"
	models "Basic/Auth-Api/Models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	var users []models.User

	initializers.DB.Find(&users)

	c.JSON(200, &users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	if n, err := strconv.Atoi(id); err == nil {
		user, err := models.GetUserByID(uint(n))
		if err == nil {
			c.JSON(200, gin.H{
				"user": &user,
			})
		}
	} else {
		c.Status(400)
	}

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
