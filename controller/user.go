package controller

import (
	"app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := service.GetUserInfo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
