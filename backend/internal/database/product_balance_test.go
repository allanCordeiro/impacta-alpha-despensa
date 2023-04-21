package database

import (
	"testing"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/database/mock_db"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateBalance(t *testing.T) {
	db := mock_db.SetupDB()
	balanceDb := NewProductBalanceDB(db)

	t.Run("Given a valid data when call save method should store data regularly", func(t *testing.T) {
		expectedProdId := "id produto"
		expectedAmount := 3
		expectedRemainingQuantity := 1

		balance := entity.NewProductBalance(expectedProdId, expectedAmount, expectedRemainingQuantity)
		err := balanceDb.Save(balance)
		assert.Nil(t, err)

		var result []entity.ProductBalance
		result, err = balanceDb.GetByProductId(balance.ProductID)
		assert.Nil(t, err)

		assert.Equal(t, 1, len(result))
		assert.Equal(t, expectedProdId, result[0].ProductID)
		assert.Equal(t, expectedAmount, result[0].DeductedAmount)
		assert.NotNil(t, result[0].DeductedDate)

	})
	mock_db.TearDown(db)
}

func TestGetBalanceList(t *testing.T) {
	db := mock_db.SetupDB()
	balanceDb := NewProductBalanceDB(db)

	t.Run("Given a valid data when call get method should shown all corresponding data", func(t *testing.T) {
		expectedProdId := "id produto"
		expectedAmount := [2]int{6, 3}
		expectedRemainingQuantity := [2]int{4, 1}

		balance := entity.NewProductBalance(expectedProdId, expectedAmount[0], expectedRemainingQuantity[0])
		err := balanceDb.Save(balance)
		assert.Nil(t, err)
		balance = entity.NewProductBalance(expectedProdId, expectedAmount[1], expectedRemainingQuantity[1])
		err = balanceDb.Save(balance)
		assert.Nil(t, err)

		var result []entity.ProductBalance
		result, err = balanceDb.GetByProductId(balance.ProductID)
		assert.Nil(t, err)

		assert.Equal(t, 2, len(result))
	})
	mock_db.TearDown(db)
}
