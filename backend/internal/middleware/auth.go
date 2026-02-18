package middleware

import (
	"net/http"
	"strings"

	"github.com/ariveratij40-lab/skia/backend/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTAuth middleware valida el token JWT
func JWTAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error: struct {
					Code    string      `json:"code"`
					Message string      `json:"message"`
					Details interface{} `json:"details,omitempty"`
				}{
					Code:    "UNAUTHORIZED",
					Message: "Authorization header required",
				},
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error: struct {
					Code    string      `json:"code"`
					Message string      `json:"message"`
					Details interface{} `json:"details,omitempty"`
				}{
					Code:    "UNAUTHORIZED",
					Message: "Invalid authorization header format",
				},
			})
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error: struct {
					Code    string      `json:"code"`
					Message string      `json:"message"`
					Details interface{} `json:"details,omitempty"`
				}{
					Code:    "UNAUTHORIZED",
					Message: "Invalid or expired token",
				},
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{
				Error: struct {
					Code    string      `json:"code"`
					Message string      `json:"message"`
					Details interface{} `json:"details,omitempty"`
				}{
					Code:    "UNAUTHORIZED",
					Message: "Invalid token claims",
				},
			})
			c.Abort()
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Set("tenant_id", claims["tenant_id"])
		c.Set("role", claims["role"])
		c.Set("email", claims["email"])

		c.Next()
	}
}

// RequireRole middleware verifica que el usuario tenga un rol específico
func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, models.ErrorResponse{
				Error: struct {
					Code    string      `json:"code"`
					Message string      `json:"message"`
					Details interface{} `json:"details,omitempty"`
				}{
					Code:    "FORBIDDEN",
					Message: "Role not found in context",
				},
			})
			c.Abort()
			return
		}

		roleStr, ok := userRole.(string)
		if !ok {
			c.JSON(http.StatusForbidden, models.ErrorResponse{
				Error: struct {
					Code    string      `json:"code"`
					Message string      `json:"message"`
					Details interface{} `json:"details,omitempty"`
				}{
					Code:    "FORBIDDEN",
					Message: "Invalid role format",
				},
			})
			c.Abort()
			return
		}

		for _, r := range roles {
			if r == roleStr {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, models.ErrorResponse{
			Error: struct {
				Code    string      `json:"code"`
				Message string      `json:"message"`
				Details interface{} `json:"details,omitempty"`
			}{
				Code:    "FORBIDDEN",
				Message: "Insufficient permissions",
			},
		})
		c.Abort()
	}
}
