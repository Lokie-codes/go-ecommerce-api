package app

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-api/internal/config"
	"go-ecommerce-api/internal/routes"
)

func Run() {
	config.LoadEnv()
	r := gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":" + config.GetEnv("API_PORT", "8080"))
}
