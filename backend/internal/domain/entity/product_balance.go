package entity

import "time"

type ProductBalance struct {
	ProductID      string
	DeductedAmount int
	DeductedDate   time.Time
}

func NewProductBalance(productId string, deductedAmount int) *ProductBalance {
	return &ProductBalance{
		ProductID:      productId,
		DeductedAmount: deductedAmount,
		DeductedDate:   time.Now(),
	}
}
