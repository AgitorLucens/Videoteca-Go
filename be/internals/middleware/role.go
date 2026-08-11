package middleware

import (
	"github.com/gin-gonic/gin"
)

func RequireRole(requireRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		r, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(403, gin.H{"error": "Role not found"})
			return
		}

		role, ok := r.(string)
		if !ok || role != requireRole {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden: insufficient role"})
			return
		}

		c.Next()
	}
}

func RequireMultipleRole(requireRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		r, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(403, gin.H{"error": "Role not found"})
			return
		}
		role, ok := r.(string)
		if !ok {
			c.AbortWithStatusJSON(403, gin.H{"error": "Invalid role type"})
			return
		}
		for _, requireRole := range requireRoles {
			if role == requireRole {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(403, gin.H{"error": "insufficient permissions"})
	}
}
