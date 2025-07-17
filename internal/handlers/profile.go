package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-ecommerce-api/internal/db"
	"go-ecommerce-api/internal/models"
)

func GetProfile(c *gin.Context) {
	userID := c.GetUint("userID")
	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, user)
}

func UpdateProfile(c *gin.Context) {
	userID := c.GetUint("userID")
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	var user models.User
	if err := db.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Password != "" {
		user.Password = req.Password
	}
	if err := db.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, user)
}
