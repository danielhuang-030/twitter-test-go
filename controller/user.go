package controller

import (
	"app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	user, err := service.GetUserInfo(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UserFollowers(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	users, err := service.GetUserFollowers(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func UserFollowing(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	users, err := service.GetUserFollowings(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
