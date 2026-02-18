package middleware

import (
	"github.com/gin-gonic/gin"
)

// RLSContext middleware establece el contexto para Row Level Security
func RLSContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantID, exists := c.Get("tenant_id")
		if !exists {
			c.Next()
			return
		}

		// El tenant_id se usará en los repositorios para establecer
		// la variable de sesión de PostgreSQL
		c.Set("rls_tenant_id", tenantID)
		c.Next()
	}
}
