package database

import (
	"testing"
	"time"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/database/mock_db"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestNewStockDb(t *testing.T) {
	db := mock_db.SetupDB()
	stockDb := NewStockDb(db)
	t.Run("Given a valid product, when calls save method then data should be stored", func(t *testing.T) {
		expectedName := "product 1"
		expectedCreationDate := time.Now()
		expectedQuantity := 20
		expectedExpirationDate := time.Now().Add(time.Hour * 24 * 10)

		product := entity.NewProduct(expectedName, expectedCreationDate, expectedQuantity, expectedExpirationDate)
		prdOk, _ := product.IsValid()
		assert.True(t, prdOk)
		err := stockDb.Save(product)
		assert.Nil(t, err)

		receivedPrd, err := stockDb.GetByID(product.ID)
		assert.Nil(t, err)
		assert.Equal(t, expectedName, receivedPrd.Name)
		assert.True(t, expectedCreationDate.Equal(receivedPrd.CreationDate))
		assert.Equal(t, expectedQuantity, receivedPrd.Quantity)
		assert.True(t, expectedExpirationDate.Equal(receivedPrd.ExpirationDate))
	})
	mock_db.TearDown(db)
}

func TestGetProduct(t *testing.T) {
	db := mock_db.SetupDB()
	stockDb := NewStockDb(db)
	t.Run("Given a valid product list when calls get product by id then should shown one product", func(t *testing.T) {
		expectedPrd1Name := "product 1"
		expectedPrd1CreationDate := time.Now()
		expectedPrd1Quantity := 20
		expectedPrd1ExpirationDate := time.Now().Add(time.Hour * 24 * 10)
		product := entity.NewProduct(
			expectedPrd1Name,
			expectedPrd1CreationDate,
			expectedPrd1Quantity,
			expectedPrd1ExpirationDate)
		err := stockDb.Save(product)
		assert.Nil(t, err)
		expectedID := product.ID

		expectedExpiredPrdName := "product 2"
		expectedExpiredPrdCreationDate := time.Now()
		expectedExpiredPrdQuantity := 20
		expectedExpiredPrdExpirationDate := time.Now().Add(time.Hour * 24 * 10)

		product = entity.NewProduct(
			expectedExpiredPrdName,
			expectedExpiredPrdCreationDate,
			expectedExpiredPrdQuantity,
			expectedExpiredPrdExpirationDate)
		err = stockDb.Save(product)
		assert.Nil(t, err)

		receivedProduct, err := stockDb.GetByID(expectedID)
		assert.Nil(t, err)
		assert.Equal(t, expectedID, receivedProduct.ID)

	})

	mock_db.TearDown(db)
}

func TestGetFullStock(t *testing.T) {
	db := mock_db.SetupDB()
	stockDb := NewStockDb(db)
	t.Run("Given a valid product, when calls save method then data should be stored", func(t *testing.T) {
		expectedPrd1Name := "product 1"
		expectedPrd1CreationDate := time.Now()
		expectedPrd1Quantity := 20
		expectedPrd1ExpirationDate := time.Now().Add(time.Hour * 24 * 10)
		product := entity.NewProduct(
			expectedPrd1Name,
			expectedPrd1CreationDate,
			expectedPrd1Quantity,
			expectedPrd1ExpirationDate)
		err := stockDb.Save(product)
		assert.Nil(t, err)
		expectedExpiredPrdName := "product 2"
		expectedExpiredPrdCreationDate := time.Now()
		expectedExpiredPrdQuantity := 20
		expectedExpiredPrdExpirationDate := time.Now().Add(time.Hour * 24 * 10)

		product = entity.NewProduct(
			expectedExpiredPrdName,
			expectedExpiredPrdCreationDate,
			expectedExpiredPrdQuantity,
			expectedExpiredPrdExpirationDate)
		err = stockDb.Save(product)
		assert.Nil(t, err)

		productList, err := stockDb.GetAllProducts()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(productList))
	})

	mock_db.TearDown(db)
}

func TestUpdateStock(t *testing.T) {
	db := mock_db.SetupDB()
	stockDb := NewStockDb(db)
	tests := []struct {
		name                      string
		quantityToDrop            int
		expectedRemainingQuantity int
		expectedError             error
	}{
		{
			name:                      "Given valid data to reduce stock when call update then should be ok",
			quantityToDrop:            2,
			expectedRemainingQuantity: 3,
			expectedError:             nil,
		},
		{
			name:                      "Given valid data to reduce stock to zero when call update then should be ok",
			quantityToDrop:            5,
			expectedRemainingQuantity: 0,
			expectedError:             nil,
		},
		{
			name:                      "Given quantity to reduce stock less than zero when call update then should return error",
			quantityToDrop:            6,
			expectedRemainingQuantity: 5,
			expectedError:             entity.ErrInsufficientStock,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedPrd := entity.NewProduct("product 2", time.Now(), 5, time.Now().Add(time.Hour*24*3))
			err := stockDb.Save(expectedPrd)
			assert.Nil(t, err)

			err = expectedPrd.UpdateQuantity(tt.quantityToDrop)
			assert.Equal(t, tt.expectedError, err)

			err = stockDb.UpdateQuantity(expectedPrd)
			assert.Nil(t, err)

			receivedPrd, err := stockDb.GetByID(expectedPrd.ID)
			assert.Nil(t, err)

			assert.Equal(t, tt.expectedRemainingQuantity, receivedPrd.Quantity)

		})
	}

	mock_db.TearDown(db)

}
