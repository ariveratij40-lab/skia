package middleware

import (
	"net/http"

	"github.com/ariveratij40-lab/skia/backend/internal/models"
	"github.com/gin-gonic/gin"
)

// ErrorHandler middleware maneja errores globales
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Error: struct {
					Code    string      `json:"code"`
					Message string      `json:"message"`
					Details interface{} `json:"details,omitempty"`
				}{
					Code:    "INTERNAL_ERROR",
					Message: err.Error(),
				},
			})
		}
	}
}
