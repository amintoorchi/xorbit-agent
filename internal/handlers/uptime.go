package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v4/host"
)

func Uptime(c *gin.Context) {
	uptime, err := host.Uptime()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"uptime":  uptime,
	})
}
