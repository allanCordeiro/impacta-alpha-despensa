package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProductBalance(t *testing.T) {
	expectedId := int64(4)
	expectedPrdId := "ProductID"
	expectedAmount := 1
	expectedDate := time.Now()

	t.Run("Given valid product balance when call NewProductBalance should return entity", func(t *testing.T) {
		prdBalance := NewProductBalance(expectedId, expectedPrdId, expectedAmount, expectedDate)
		assert.NotNil(t, prdBalance)
		assert.Equal(t, expectedId, prdBalance.ID)
		assert.Equal(t, expectedPrdId, prdBalance.ProductID)
		assert.Equal(t, expectedAmount, prdBalance.DeductedAmount)
		assert.Equal(t, expectedDate, prdBalance.DeductedDate)
	})
}
