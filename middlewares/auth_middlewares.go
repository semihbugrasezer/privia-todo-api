package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/semihbugrasezer/privia-todo-api/utils"
)

// AuthMiddleware JWT doğrulama işlemlerini yapar
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.RespondWithError(c, http.StatusUnauthorized, "Authorization header bulunamadı")
			c.Abort()
			return
		}

		// Bearer token şeklinde olup olmadığını kontrol et
		tokenString := strings.Split(authHeader, "Bearer ")
		if len(tokenString) != 2 {
			utils.RespondWithError(c, http.StatusUnauthorized, "Geçersiz Authorization header formatı")
			c.Abort()
			return
		}

		// JWT tokeni doğrula
		token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
			// Tokenin doğrulanma metodunun doğru olup olmadığını kontrol et
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("Geçersiz token doğrulama metodu", jwt.ValidationErrorSignatureInvalid)
			}
			// Buraya gizli anahtarınızı girin
			return []byte("secret"), nil
		})

		if err != nil {
			utils.RespondWithError(c, http.StatusUnauthorized, "Geçersiz token")
			c.Abort()
			return
		}

		// Token geçerli mi kontrol et
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userID", claims["userID"])
		} else {
			utils.RespondWithError(c, http.StatusUnauthorized, "Geçersiz token")
			c.Abort()
			return
		}

		c.Next()
	}
}
