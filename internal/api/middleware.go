package api

import (
	"crypto/subtle"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestToken := c.GetHeader("X-Orbit-Token")

		if subtle.ConstantTimeCompare([]byte(requestToken), []byte(token)) != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized",
			})
			return
		}

		c.Next()
	}
}
