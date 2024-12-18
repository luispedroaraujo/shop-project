package models

var TestProducts = []Product{
	{
		SKU:      "000001",
		Name:     "Product",
		Category: "sneakers",
		Price:    10000,
	},
	{
		SKU:      "000003",
		Name:     "Special SKU Product",
		Category: "sandals",
		Price:    20000,
	},
	{
		SKU:      "000002",
		Name:     "BV Lean leather ankle boots",
		Category: "boots",
		Price:    99000,
	},
	{
		SKU:      "000003",
		Name:     "Boots with SKU Discount",
		Category: "boots",
		Price:    40000,
	},
}

var TestProductsWithDiscount = []ProductWithDiscount{
	{
		Product: TestProducts[0],
		Price: Price{
			Original:           10000,
			Final:              10000,
			Currency:           "EUR",
			DiscountPercentage: "",
		},
	},
	{
		Product: TestProducts[1],
		Price: Price{
			Original:           20000,
			Final:              17000,
			Currency:           "EUR",
			DiscountPercentage: "15%",
		},
	},
	{
		Product: TestProducts[2],
		Price: Price{
			Original:           99000,
			Final:              69300,
			Currency:           "EUR",
			DiscountPercentage: "30%",
		},
	},
	{
		Product: TestProducts[3],
		Price: Price{
			Original:           40000,
			Final:              28000,
			Currency:           "EUR",
			DiscountPercentage: "30%",
		},
	},
}
