package main

import (
	initializers "Basic/Auth-Api/Initializers"
	models "Basic/Auth-Api/Models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(models.User{}, models.Role{})
}
