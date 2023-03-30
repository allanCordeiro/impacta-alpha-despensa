package entity

import "time"

type ProductBalance struct {
	ID             int64
	ProductID      string
	DeductedAmount int
	DeductedDate   time.Time
}

func NewProductBalance(id int64, productId string, deductedAmount int, deductedDate time.Time) *ProductBalance {
	return &ProductBalance{
		ID:             id,
		ProductID:      productId,
		DeductedAmount: deductedAmount,
		DeductedDate:   deductedDate,
	}
}
