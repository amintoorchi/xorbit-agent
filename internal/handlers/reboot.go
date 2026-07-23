package handlers

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func Reboot(c *gin.Context) {
	err := exec.Command("systemctl", "reboot").Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Server is rebooting.",
	})
}
