package dto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func WriteErrorResponse(c *gin.Context, statusCode int, errMsg string) {
	c.JSON(statusCode, ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: errMsg,
	})
}
