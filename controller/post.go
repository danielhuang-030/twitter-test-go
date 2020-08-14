package controller

import (
	"app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreatePostRequest struct {
	Content string `json:"content" binding:"required"`
}

type EditPostRequest struct {
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

func ShowPost(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	post, err := service.GetPost(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

func EditPost(c *gin.Context) {
	var requestData EditPostRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var data map[string]interface{}
	data = make(map[string]interface{})
	data["content"] = requestData.Content
	post, err := service.EditPost(uint(id), data, c.MustGet("currentUser"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	err := service.DeletePost(uint(id), c.MustGet("currentUser"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted post!",
	})
}
