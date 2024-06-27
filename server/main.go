package main

import (
	"example/server/controllers"
	"example/server/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/test", controllers.TestModel)

	// Users
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetDetails)
	//r.POST("/users/:id/follow")
	//r.POST("/users/:id/follow")
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)

	err := r.Run()
	if err != nil {
		return
	}
}
