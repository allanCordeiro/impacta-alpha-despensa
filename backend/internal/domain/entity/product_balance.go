package entity

import "time"

type ProductBalance struct {
	ProductID         string
	DeductedAmount    int
	RemainingQuantity int
	DeductedDate      time.Time
}

func NewProductBalance(productId string, deductedAmount int, remainingQuantity int) *ProductBalance {
	return &ProductBalance{
		ProductID:         productId,
		DeductedAmount:    deductedAmount,
		RemainingQuantity: remainingQuantity,
		DeductedDate:      time.Now(),
	}
}
