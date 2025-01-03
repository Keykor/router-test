package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// FakeAuthMiddleware is a middleware that checks for the presence of an Authorization header with a Bearer token
// it's a fake authentication used for testing purposes
func FakeAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid"})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		c.Set("token", token)

		c.Next()
	}
}
