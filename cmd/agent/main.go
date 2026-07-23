package main

import (
	"go-agent/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api.RegisterRoutes(router)

	router.Run(":42100")
}
