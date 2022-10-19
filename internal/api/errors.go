package api

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Message string `json:"error"`
}

func writeErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, ErrorResponse{
		Message: message,
	})
}
