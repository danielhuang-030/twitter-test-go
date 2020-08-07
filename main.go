package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	api := router.Group("/api")
	api.GET("/", index)

	godotenv.Load()
	router.Run(":" + os.Getenv("APP_PORT"))
}

func index(c *gin.Context) {
	statusCode := http.StatusOK
	c.JSON(statusCode, gin.H{
		"data": "Hello, GIN",
	})
}
