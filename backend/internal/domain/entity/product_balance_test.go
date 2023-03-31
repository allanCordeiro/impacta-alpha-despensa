package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductBalance(t *testing.T) {
	expectedPrdId := "ProductID"
	expectedAmount := 1

	t.Run("Given valid product balance when call NewProductBalance should return entity", func(t *testing.T) {
		prdBalance := NewProductBalance(expectedPrdId, expectedAmount)
		assert.NotNil(t, prdBalance)
		assert.Equal(t, expectedPrdId, prdBalance.ProductID)
		assert.Equal(t, expectedAmount, prdBalance.DeductedAmount)
		assert.NotNil(t, prdBalance.DeductedDate)
	})
}
