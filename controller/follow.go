package controller

import (
	"app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateFollowRequest struct {
	UserID uint `json:"user_id" binding:"required"`
}

func CreateFollow(c *gin.Context) {
	var requestData CreateFollowRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := service.CreateFollower(requestData.UserID, c.MustGet("currentUser"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully followed user!",
	})
}

func DeleteFollow(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	err := service.DeleteFollower(uint(id), c.MustGet("currentUser"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully unfollowed user!",
	})
}
