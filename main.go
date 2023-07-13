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
	initializers.ConnectToRedis()
}

func main() {

	r := gin.Default()

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)

	authenticated := r.Group("/api")
	authenticated.Use(middlewares.JwtAuthMiddleware())
	authenticated.GET("/userinfo", controllers.CurrentUser)
	authenticated.GET("/users", controllers.ListUsers)
	authenticated.GET("/users/:id", controllers.GetUser)
	authenticated.PUT("/profile-img", controllers.UpdateProfileImg)
	authenticated.DELETE("/profile-img", controllers.RemoveProfileImg)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.Use(middlewares.RoleBaseAuthmiddleware("admin"))
	protected.PUT("/users/:id", controllers.UpdateUser)
	protected.DELETE("/users/:id", controllers.DeleteUser)

	protected.POST("/roles", controllers.PostRole)
	protected.GET("/roles", controllers.ListRoles)
	protected.GET("/roles/:id", controllers.GetRole)

	r.Run()

	// arr := []int{1, 2, 2, 3, 4, 5, 6, 7, 8, 5, 9, 9, 10}
	// filt := []int{}
	// dupArr := []int{}
	// for _, v := range arr {
	// 	for _, v2 := range filt {
	// 		if v == v2 {
	// 			dupArr = append(dupArr, v)
	// 			break
	// 		}
	// 	}
	// 	filt = append(filt, v)

	// }

	// for _, vf := range dupArr {
	// 	fmt.Println(vf)
	// }
}
