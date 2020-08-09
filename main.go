package main

import (
	"app/controller"
	_ "app/init"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/", controller.Index)
	api.POST("/signup", controller.Signup)
	api.POST("/login", controller.Login)
	router.Run(":" + os.Getenv("APP_PORT"))
}
