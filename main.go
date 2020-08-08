package main

import (
	"app/controller"
	_ "app/init"
	_ "app/model"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/", controller.Index)

	router.Run(":" + os.Getenv("APP_PORT"))
}
