package main

import (
	"example/server/controllers"
	"example/server/libs"
	"example/server/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	authorized := r.Group("/")

	authorized.Use(libs.IsAuthenticated())

	r.GET("/test", controllers.TestModel)

	// Users
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.GetDetails)
	//authorized.POST("/users/:id/follow")
	//authorized.POST("/users/:id/follow")
	r.POST("/login", controllers.Login)
	authorized.POST("/logout", controllers.Logout)

	err := r.Run()
	if err != nil {
		return
	}
}
