package controller

import (
	"app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePostRequest struct {
	Content string `json:"content" binding:"required"`
}

func CreatePost(c *gin.Context) {
	var requestData CreatePostRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	var data map[string]interface{}
	data = make(map[string]interface{})
	data["content"] = requestData.Content
	post, err := service.CreatePost(data, c.MustGet("currentUser"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, post)
}
