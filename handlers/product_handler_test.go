package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"shop-api/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Setup DB and router
func setupTest(t *testing.T) *gin.Engine {
	gin.SetMode(gin.TestMode)

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to in-memory database: %v", err)
	}

	if err := db.AutoMigrate(&models.Product{}); err != nil {
		t.Fatalf("Failed to migrate database schema: %v", err)
	}

	if err := db.Create(&models.TestProducts).Error; err != nil {
		t.Fatalf("Failed to add test data: %v", err)
	}

	router := gin.Default()
	router.GET("/products", func(c *gin.Context) {
		GetProducts(c, db)
	})

	return router
}

func TestGetAllProducts(t *testing.T) {
	router := setupTest(t)

	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.ProductWithDiscount
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Len(t, response, 4)
	assert.Equal(t, models.TestProductsWithDiscount, response)
}

func TestGetProductsWithCategoryAndPriceLessThan(t *testing.T) {
	router := setupTest(t)

	req, _ := http.NewRequest(http.MethodGet, "/products?category=boots&priceLessThan=40000", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.ProductWithDiscount
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Len(t, response, 1)
	assert.Equal(t, models.TestProductsWithDiscount[3], response[0])
}
