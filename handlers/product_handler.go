package handlers

import (
	"net/http"
	"shop-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProducts(c *gin.Context, db *gorm.DB) {
	var products []models.Product
	category := c.DefaultQuery("category", "")
	priceLessThanStr := c.DefaultQuery("priceLessThan", "")
	query := db.Model(&models.Product{})

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if priceLessThanStr != "" {
		priceLessThan, err := strconv.Atoi(priceLessThanStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid priceLessThan value"})
			return
		}
		query = query.Where("price <= ?", priceLessThan)
	}

	if err := query.Limit(5).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	productsWithDiscount := []models.ProductWithDiscount{}
	for _, product := range products {
		p := models.ApplyDiscount(product)
		productsWithDiscount = append(productsWithDiscount, p)
	}

	c.JSON(http.StatusOK, productsWithDiscount)
}
