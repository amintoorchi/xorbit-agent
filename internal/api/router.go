package api

import (
	"go-agent/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Go Agent API",
			"success": true,
		})
	})

	router.GET("/pair", handlers.Pair)
	router.GET("/up", handlers.Health)
	router.GET("/exec/shutdown", handlers.Shutdown)
	router.GET("/exec/reboot", handlers.Reboot)
	router.GET("/metrics/os/uptime", handlers.Uptime)
	router.GET("/metrics/os/info", handlers.Info)
	router.GET("/metrics/hw/cpu", handlers.CPU)
}
