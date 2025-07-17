package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-ecommerce-api/internal/db"
	"go-ecommerce-api/internal/models"
)

func CreateOrder(c *gin.Context) {
	userID := c.GetUint("userID")
	var cartItems []models.CartItem
	if err := db.DB.Where("user_id = ?", userID).Preload("Product").Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart"})
		return
	}
	if len(cartItems) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cart is empty"})
		return
	}
	var total float64
	var orderItems []models.OrderItem
	for _, item := range cartItems {
		orderItems = append(orderItems, models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Product:   item.Product,
		})
		total += item.Product.Price * float64(item.Quantity)
	}
	order := models.Order{UserID: userID, Items: orderItems, Total: total}
	if err := db.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}
	// Clear cart
	db.DB.Where("user_id = ?", userID).Delete(&models.CartItem{})
	c.JSON(http.StatusCreated, order)
}

func GetOrders(c *gin.Context) {
	userID := c.GetUint("userID")
	var orders []models.Order
	if err := db.DB.Where("user_id = ?", userID).Preload("Items.Product").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func GetOrder(c *gin.Context) {
	userID := c.GetUint("userID")
	id := c.Param("id")
	var order models.Order
	if err := db.DB.Where("id = ? AND user_id = ?", id, userID).Preload("Items.Product").First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}
