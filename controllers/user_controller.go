package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/semihbugrasezer/privia-todo-api/mock"
	"github.com/semihbugrasezer/privia-todo-api/models"
	"github.com/semihbugrasezer/privia-todo-api/utils"
)

// Login handles user login and generates a JWT token
func Login(c *gin.Context) {
	var credentials models.User
	if err := c.ShouldBindJSON(&credentials); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid login credentials")
		return
	}

	// Kullanıcı doğrulama (mock data ile karşılaştırma)
	for _, user := range mock.Users {
		if user.Username == credentials.Username && user.Password == credentials.Password {
			// Token oluşturma
			expirationTime := time.Now().Add(24 * time.Hour) // Token 24 saat geçerli olacak
			claims := &utils.JWTClaims{
				UserID:   user.ID,
				UserType: user.Type,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(utils.JwtKey)
			if err != nil {
				utils.RespondWithError(c, http.StatusInternalServerError, "Token oluşturulamadı")
				return
			}

			utils.RespondWithSuccess(c, http.StatusOK, "Giriş başarılı", gin.H{"token": tokenString})
			return
		}
	}

	utils.RespondWithError(c, http.StatusUnauthorized, "Geçersiz kullanıcı adı veya parola")
}
