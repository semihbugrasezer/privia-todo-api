package controllers

import (
	"net/http"
	"privia-staj-backend-todo/mock"
	"privia-staj-backend-todo/models"
	"privia-staj-backend-todo/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var requestBody models.LoginRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utils.BadRequest(c, "Invalid input", err)
		return
	}

	var user models.User
	if err := mock.DB.Where("username = ?", requestBody.Username).First(&user).Error; err != nil {
		utils.Unauthorized(c, "Incorrect username or password: username not found", err)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password)); err != nil {
		utils.Unauthorized(c, "Incorrect username or password: password mismatch", err)
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.UserType)
	if err != nil {
		utils.InternalServerError(c, "Could not generate token", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
