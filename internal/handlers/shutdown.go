package handlers

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func Shutdown(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Server is shutting down.",
	})

	go exec.Command("systemctl", "poweroff").Run()
}
