package main

import (
	"app/controller"
	_ "app/init"
	"app/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/", controller.Index)
	api.POST("/signup", controller.Signup)
	api.POST("/login", controller.Login)

	api.Use(middleware.JWT())
	{
		api.GET("/logout", controller.Logout)
		api.GET("/users/:id/info", controller.UserInfo)
		api.GET("/users/:id/followers", controller.UserFollowers)
		api.GET("/users/:id/following", controller.UserFollowing)
		api.PATCH("/following/:id", controller.CreateFollow)
		api.DELETE("/following/:id", controller.DeleteFollow)
		api.GET("/posts/:id", controller.ShowPost)
		api.POST("/posts", controller.CreatePost)
		api.PUT("/posts/:id", controller.EditPost)
		api.DELETE("/posts/:id", controller.DeletePost)
	}
	router.Run(":" + os.Getenv("APP_PORT"))
}
