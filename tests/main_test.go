package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"go-ecommerce-api/internal/app"
)

func TestHealth(t *testing.T) {
	w := httptest.NewRecorder()
	r := setupRouter()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.ServeHTTP(w, httptest.NewRequest("GET", "/health", nil))
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func setupRouter() *gin.Engine {
	app.Run()
	return gin.Default()
}
