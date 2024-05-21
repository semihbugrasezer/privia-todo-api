package middlewares

import (
	"net/http"
	"privia-staj-backend-todo/utils"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// AuthMiddleware JWT doğrulama işlemlerini yapar
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.RespondError(c, http.StatusUnauthorized, "Authorization header bulunamadı")
			c.Abort()
			return
		}

		// Bearer token şeklinde olup olmadığını kontrol et
		tokenString := strings.Split(authHeader, "Bearer ")
		if len(tokenString) != 2 {
			utils.RespondError(c, http.StatusUnauthorized, "Geçersiz Authorization header formatı")
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
			utils.RespondError(c, http.StatusUnauthorized, "Geçersiz token")
			c.Abort()
			return
		}

		// Token geçerli mi kontrol et
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userID", claims["userID"])
		} else {
			utils.RespondError(c, http.StatusUnauthorized, "Geçersiz token")
			c.Abort()
			return
		}

		c.Next()
	}
}
