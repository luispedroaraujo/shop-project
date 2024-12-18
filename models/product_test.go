package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyDiscount(t *testing.T) {
	tests := []struct {
		name     string
		product  Product
		expected ProductWithDiscount
	}{
		{
			name:     "No discount",
			product:  TestProducts[0],
			expected: TestProductsWithDiscount[0],
		},
		{
			name:     "15% discount for SKU 000003",
			product:  TestProducts[1],
			expected: TestProductsWithDiscount[1],
		},
		{
			name:     "30% discount for boots category",
			product:  TestProducts[2],
			expected: TestProductsWithDiscount[2],
		},
		{
			name:     "30% boots discount takes precedence over SKU discount",
			product:  TestProducts[3],
			expected: TestProductsWithDiscount[3],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ApplyDiscount(tt.product)
			assert.Equal(t, tt.expected, result)
		})
	}
}
