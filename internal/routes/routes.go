package routes

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-api/internal/handlers"
)

func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api/v1")

	api.POST("/register", handlers.Register)
	api.POST("/login", handlers.Login)

	api.GET("/products", handlers.GetProducts)
	api.GET("/products/:id", handlers.GetProduct)
	api.POST("/products", handlers.AuthMiddleware(true), handlers.CreateProduct)
	api.PUT("/products/:id", handlers.AuthMiddleware(true), handlers.UpdateProduct)
	api.DELETE("/products/:id", handlers.AuthMiddleware(true), handlers.DeleteProduct)

	api.GET("/cart", handlers.AuthMiddleware(false), handlers.GetCart)
	api.POST("/cart/add", handlers.AuthMiddleware(false), handlers.AddToCart)
	api.DELETE("/cart/remove/:itemId", handlers.AuthMiddleware(false), handlers.RemoveFromCart)

	api.POST("/orders", handlers.AuthMiddleware(false), handlers.CreateOrder)
	api.GET("/orders", handlers.AuthMiddleware(false), handlers.GetOrders)
	api.GET("/orders/:id", handlers.AuthMiddleware(false), handlers.GetOrder)

	api.GET("/profile", handlers.AuthMiddleware(false), handlers.GetProfile)
	api.PUT("/profile", handlers.AuthMiddleware(false), handlers.UpdateProfile)

	r.GET("/health", handlers.Health)
}
