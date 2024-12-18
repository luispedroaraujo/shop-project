package routes

import (
	"shop-api/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterProductRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/products", func(c *gin.Context) { handlers.GetProducts(c, db) })
}
