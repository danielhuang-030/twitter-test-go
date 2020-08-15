package controller

import (
	"app/middleware"
	"app/service"
	"encoding/json"
	"net/http"
	"strings"

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

type userResponse struct {
}

func Signup(c *gin.Context) {
	var requestData SignupRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
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

	c.JSON(http.StatusOK, user)
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

	// response with token
	userJSON, _ := json.Marshal(user)
	userMap := make(map[string]interface{})
	json.Unmarshal(userJSON, &userMap)
	userMap["token"] = token

	c.JSON(http.StatusOK, userMap)
}

func Logout(c *gin.Context) {
	splitToken := strings.Split(c.GetHeader("Authorization"), "Bearer ")
	token := splitToken[1]
	claims, _ := middleware.ParseToken(token)
	err := service.Logout(token, claims.ExpiresAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}
