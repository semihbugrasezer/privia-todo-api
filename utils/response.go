package utils

import (
	"github.com/gin-gonic/gin"
)

// Standart bir hata yanıtı gönderir
func RespondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

// Başarılı bir JSON yanıtı gönderir
func RespondWithSuccess(c *gin.Context, code int, message interface{}, data interface{}) {
	c.JSON(code, gin.H{"message": message, "data": data})
}

// JSON yanıtı gönderir (başarılı veya başarısız olabilir)
func RespondJSON(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, payload)
}
