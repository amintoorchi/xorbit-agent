package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"method":  c.Request.Method,
		"host":    c.Request.Host, // localhost:8080
		"url":     c.Request.URL.Path,
		"message": "Server is successfully running",
	})
}
