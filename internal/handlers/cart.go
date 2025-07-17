package handlers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"go-ecommerce-api/internal/db"
	"go-ecommerce-api/internal/models"
)

func GetCart(c *gin.Context) {
	userID := c.GetUint("userID")
	var items []models.CartItem
	if err := db.DB.Where("user_id = ?", userID).Preload("Product").Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart"})
		return
	}
	c.JSON(http.StatusOK, items)
}

func AddToCart(c *gin.Context) {
	userID := c.GetUint("userID")
	var req struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Quantity < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	item := models.CartItem{UserID: userID, ProductID: req.ProductID, Quantity: req.Quantity}
	if err := db.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to cart"})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func RemoveFromCart(c *gin.Context) {
	userID := c.GetUint("userID")
	itemId := c.Param("itemId")
	var item models.CartItem
	if err := db.DB.Where("id = ? AND user_id = ?", itemId, userID).First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	if err := db.DB.Delete(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove item"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Item removed"})
}
