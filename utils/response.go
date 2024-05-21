package utils

import "github.com/gin-gonic/gin"

// RespondError sends a standard error response
func RespondError(c *gin.Context, code int, message string) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

// RespondSuccess sends a success JSON response
func RespondSuccess(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{"message": message, "data": data})
}

// RespondJSON sends a JSON response (can be success or error)
func RespondJSON(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, payload)
}
