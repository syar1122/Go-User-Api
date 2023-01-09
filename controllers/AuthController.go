package controllers

import (
	initializers "Basic/Auth-Api/Initializers"
	models "Basic/Auth-Api/Models"
	token "Basic/Auth-Api/Token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {

	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)
	u.PrepareGive()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=15"`
}

func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheck(u.Username, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

type RegisterInput struct {
	FullName string      `json:"fullName" binding:"required"`
	Password string      `json:"password" binding:"required,min=8,max=15"`
	Username string      `json:"userName" binding:"required,alpha"`
	RoleId   int         `json:"roleId" binding:"required"`
	Role     models.Role `json:"role"`
	Status   int8        `json:"status" binding:"required,min=1,max=2"`
}

func Register(c *gin.Context) {

	var body RegisterInput

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isUserExist := models.IsExist(body.Username)
	if isUserExist {
		c.JSON(400, gin.H{"message": "UserName has been taken!!!"})
		return
	}

	user := models.User{FullName: body.FullName, Password: body.Password, Username: body.Username, RoleId: body.RoleId, Status: body.Status}
	result := initializers.DB.Create(&user).Preload("Role").First(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "user": &user})

}
