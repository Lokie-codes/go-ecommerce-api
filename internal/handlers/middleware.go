package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-ecommerce-api/internal/utils"
)

func AuthMiddleware(adminOnly bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, isAdmin, err := utils.ValidateJWT(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		if adminOnly && !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}
		c.Set("userID", userID)
		c.Set("isAdmin", isAdmin)
		c.Next()
	}
}
