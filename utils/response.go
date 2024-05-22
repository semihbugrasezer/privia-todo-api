package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response is a standard response format
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse is a standard error response format
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// SendResponse sends a JSON response with a custom message and data
func SendResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

// SendErrorResponse sends a JSON error response with a custom message and error details
func SendErrorResponse(c *gin.Context, status int, message string, err error) {
	c.JSON(status, ErrorResponse{
		Status:  status,
		Message: message,
		Error:   err.Error(),
	})
}

// OK sends a 200 OK response
func OK(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"message": message, "data": data})
}

// Created sends a 201 Created response
func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{"message": message, "data": data})
}

// BadRequest sends a 400 Bad Request response
func BadRequest(c *gin.Context, message string, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"message": message, "error": err.Error()})
}

// Unauthorized sends a 401 Unauthorized response
func Unauthorized(c *gin.Context, message string, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{"message": message, "error": err.Error()})
}

// Forbidden sends a 403 Forbidden response
func Forbidden(c *gin.Context, message string, err error) {
	c.JSON(http.StatusForbidden, gin.H{"message": message, "error": err.Error()})
}

// NotFound sends a 404 Not Found response
func NotFound(c *gin.Context, message string, err error) {
	c.JSON(http.StatusNotFound, gin.H{"message": message, "error": err.Error()})
}

// InternalServerError sends a 500 Internal Server Error response
func InternalServerError(c *gin.Context, message string, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": message, "error": err.Error()})
}

// RespondJSON sends a JSON response with the given status code
func RespondJSON(c *gin.Context, statusCode int, payload interface{}) {
	c.JSON(statusCode, payload)
}

// RespondError sends a JSON error response with the given status code
func RespondError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}
