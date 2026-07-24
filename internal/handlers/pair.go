package handlers

import (
	"net/http"

	"go-agent/internal/system"

	"github.com/gin-gonic/gin"
)

func Pair(c *gin.Context) {
	info, err := system.GetInfo()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    info,
	})
}
