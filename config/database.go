package config

import (
	"encoding/json"
	"log"
	"os"
	"shop-api/models"
	"shop-api/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("shop.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	if err := db.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}

	loadProducts(db)

	return db
}

// Load products from file if DB is empty
func loadProducts(db *gorm.DB) {
	res := db.Take(&models.Product{})
	if res.RowsAffected == 0 {
		fileData, err := os.ReadFile(utils.GetEnv("PRODUCTS_FILE", "products.json"))
		if err != nil {
			log.Fatalf("Error reading JSON file: %v", err)
		}

		var products []models.Product
		err = json.Unmarshal(fileData, &products)
		if err != nil {
			log.Fatalf("Error unmarshaling JSON: %v", err)
		}

		if err := db.Create(&products).Error; err != nil {
			log.Fatalf("Error inserting data into the database: %v", err)
		}
	}
}
