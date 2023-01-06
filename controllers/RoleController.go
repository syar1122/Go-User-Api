package controllers

import (
	initializers "Basic/Auth-Api/Initializers"
	models "Basic/Auth-Api/Models"
	"log"

	"github.com/gin-gonic/gin"
)

func ListRoles(c *gin.Context) {
	var roles []models.Role
	initializers.DB.Find(&roles)
	c.JSON(200, roles)
}

func GetRole(c *gin.Context) {
	id := c.Param("id")

	var role models.Role
	initializers.DB.First(&role, id)
	c.JSON(200, role)
}

func PostRole(c *gin.Context) {

	var body struct {
		Name string
	}
	c.Bind(&body)

	role := models.Role{Name: body.Name}
	result := initializers.DB.Create(&role)

	if result.Error != nil {
		c.Status(400)
		log.Fatal(result.Error)
		return
	}

	c.JSON(200, gin.H{
		"Role": &role,
	})

}
