package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v4/cpu"
)

func CPU(c *gin.Context) {
	logical, err := cpu.Counts(true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	physical, err := cpu.Counts(false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	info, err := cpu.Info()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":        true,
		"physical_cores": physical,
		"logical_cores":  logical,
		"model":          info[0].ModelName,
		// "mhz":            info[0].Mhz,
	})
}
