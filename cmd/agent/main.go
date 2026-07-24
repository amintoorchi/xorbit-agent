package main

import (
	"log"

	"go-agent/internal/api"
	"go-agent/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatal(err)
	}

	router.Use(api.AuthMiddleware(cfg.Token))

	api.RegisterRoutes(router)

	if err := router.Run(":42100"); err != nil {
		log.Fatal(err)
	}
}
