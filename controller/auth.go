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

func Signup(c *gin.Context) {
	var requestData SignupRequest
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.BindJSON(&requestData)
	var data map[string]interface{}
	data = make(map[string]interface{})
	data["name"] = requestData.Name
	data["email"] = requestData.Email
	data["password"] = requestData.Password
	user := service.CreateUser(data)

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}
