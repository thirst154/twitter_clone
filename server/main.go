package main

import (
	"example/server/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/test", controllers.TestModel)

	err := r.Run()
	if err != nil {
		return
	}
}
