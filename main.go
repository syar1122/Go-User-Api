package main

import (
	initializers "Basic/Auth-Api/Initializers"
	"Basic/Auth-Api/controllers"
	"Basic/Auth-Api/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.POST("/roles", controllers.PostRole)
	r.GET("/roles", controllers.ListRoles)
	r.GET("/roles/:id", controllers.GetRole)

	r.POST("/users", controllers.PostUser)
	r.GET("/users", controllers.ListUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/userinfo", controllers.CurrentUser)

	r.Run()
}
