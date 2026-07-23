package main

import (
	"log"

	"go-agent/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatal(err)
	}

	api.RegisterRoutes(router)

	if err := router.Run(":42100"); err != nil {
		log.Fatal(err)
	}
}
