package controllers

import (
	"net/http"
	"privia-staj-backend-todo/mock"
	"privia-staj-backend-todo/models"
	"privia-staj-backend-todo/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Login handles user login and generates a JWT token
func Login(c *gin.Context) {
	var credentials models.User
	if err := c.ShouldBindJSON(&credentials); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Geçersiz giriş bilgileri")
		return
	}

	// Kullanıcı doğrulama (mock data ile karşılaştırma)
	for _, user := range mock.Users {
		if user.Username == credentials.Username && user.Password == credentials.Password {
			// Token oluşturma
			expirationTime := time.Now().Add(24 * time.Hour) // Token 24 saat geçerli olacak
			claims := &utils.JWTClaims{
				UserID:   user.ID,
				UserType: user.UserType, // UserType alanı düzeltildi
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(utils.JwtKey)
			if err != nil {
				utils.RespondError(c, http.StatusInternalServerError, "Token oluşturulamadı")
				return
			}
			utils.RespondSuccess(c, http.StatusOK, "Giriş başarılı", gin.H{"token": tokenString})
			return
		}
	}
	utils.RespondError(c, http.StatusUnauthorized, "Geçersiz kullanıcı adı veya parola")
}
