// response.go

package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response represents a standard response format
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse generates an error response with the given status code and message
func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{
		Status:  statusCode,
		Message: message,
	})
}

// SuccessResponse generates a success response with the given data
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	})
}

// NotFoundResponse generates a not found response with the given message
func NotFoundResponse(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Response{
		Status:  http.StatusNotFound,
		Message: message,
	})
}
