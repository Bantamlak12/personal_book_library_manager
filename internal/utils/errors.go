package utils

import (
	"github.com/Bantamlak12/personal_book_library_manager/internal/models"
	"github.com/gin-gonic/gin"
)

func NewErrorResponse(c *gin.Context, statusCode int, code, message, details string) {
	response := models.ErrorResponse{
		Status:  statusCode,
		Code:    code,
		Message: message,
		Details: details,
	}

	c.JSON(statusCode, response)
}
