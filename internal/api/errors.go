package api

import (
	"github.com/gin-gonic/gin"

	"github.com/prybintsev/stakefish/internal/models"
)

func writeErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, models.ErrorResponse{
		Message: message,
	})
}
