package controller

import (
	"app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupRequest struct {
	Name                 string `json:"name" binding:"required"`
	Email                string `json:"email" binding:"required,email"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Signup(c *gin.Context) {
	var requestData SignupRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var data map[string]interface{}
	data = make(map[string]interface{})
	data["name"] = requestData.Name
	data["email"] = requestData.Email
	data["password"] = requestData.Password
	user, err := service.CreateUser(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}

func Login(c *gin.Context) {
	var requestData LoginRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	var data map[string]interface{}
	data = make(map[string]interface{})
	data["email"] = requestData.Email
	data["password"] = requestData.Password
	user, token, err := service.Attempt(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  user,
		"token": token,
	})
}
