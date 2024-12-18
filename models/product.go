package models

import (
	"encoding/json"
)

type Product struct {
	ID       uint   `gorm:"primaryKey" json:"-"`
	SKU      string `gorm:"not null" json:"sku"`
	Name     string `gorm:"not null" json:"name"`
	Category string `gorm:"not null" json:"category"`
	Price    uint   `gorm:"not null;check:price >= 0" json:"price"`
}

type ProductWithDiscount struct {
	Product Product
	Price   Price `json:"price"`
}

func (p *ProductWithDiscount) UnmarshalJSON(data []byte) error {
	// Define an intermediate structure that matches the JSON
	type Temp struct {
		SKU      string `json:"sku"`
		Name     string `json:"name"`
		Category string `json:"category"`
		Price    Price  `json:"price"`
	}

	var temp Temp
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Map the intermediate structure to the final ProductWithDiscount struct
	p.Product = Product{
		SKU:      temp.SKU,
		Name:     temp.Name,
		Category: temp.Category,
		Price:    temp.Price.Original,
	}
	p.Price = temp.Price

	return nil
}

func (p ProductWithDiscount) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SKU      string `json:"sku"`
		Name     string `json:"name"`
		Category string `json:"category"`
		Price    Price  `json:"price"`
	}{
		SKU:      p.Product.SKU,
		Name:     p.Product.Name,
		Category: p.Product.Category,
		Price:    p.Price,
	})
}

type Price struct {
	Original           uint       `json:"original"`
	Final              uint       `json:"final"`
	DiscountPercentage NullString `json:"discount_percentage"`
	Currency           string     `json:"currency"`
}

// Handling empty strings as null
type NullString string

func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns == "" {
		return []byte("null"), nil
	}
	return json.Marshal(string(ns))
}

func ApplyDiscount(p Product) ProductWithDiscount {
	productResponse := ProductWithDiscount{
		Product: p,
		Price: Price{
			Original: p.Price,
			Final:    p.Price,
			Currency: "EUR",
		},
	}

	// 15% discount for SKU "000003"
	if p.SKU == "000003" {
		productResponse.Price.Final = uint(float64(p.Price) * 0.85)
		productResponse.Price.DiscountPercentage = "15%"
	}

	// 30% discount for "boots" category
	if p.Category == "boots" {
		productResponse.Price.Final = uint(float64(p.Price) * 0.7)
		productResponse.Price.DiscountPercentage = "30%"
	}

	return productResponse
}
