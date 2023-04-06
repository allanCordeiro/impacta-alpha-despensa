package usecase

import (
	"context"
	"database/sql"

	"testing"
	"time"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/database"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/database/mock_db"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/AllanCordeiro/impacta-alpha-despensa/pkg/uow"
	"github.com/stretchr/testify/assert"
)

// db would be a sqlite memory
// the main idea is to test product and stockgateway update and product balace create instance at the same time.
// postgres should be tested on e2e tests
func TestProductBalanceUpdateIntegrationTest(t *testing.T) {
	product := *entity.NewProduct("product 1", time.Now(), 5, time.Now().Add(time.Hour*24*5))

	tests := []struct {
		name                string
		expectedStock       entity.Product
		firstReduce         UpdateProductInput
		secondReduce        UpdateProductInput
		expectedFirstResult *UpdateProductOutput
		expectedSecResult   *UpdateProductOutput
		expectedError       error
	}{
		{
			name:                "Given a product balance amout update when execute use case should return a valid product balance",
			expectedStock:       product,
			firstReduce:         UpdateProductInput{ProductID: product.ID, Quantity: 2},
			secondReduce:        UpdateProductInput{},
			expectedFirstResult: &UpdateProductOutput{RemainingQuantity: 3},
			expectedSecResult:   nil,
			expectedError:       nil,
		},
		{
			name:                "Given two product balances amount update when execute use case should return a valid product balance",
			expectedStock:       product,
			firstReduce:         UpdateProductInput{ProductID: product.ID, Quantity: 3},
			secondReduce:        UpdateProductInput{ProductID: product.ID, Quantity: 2},
			expectedFirstResult: &UpdateProductOutput{RemainingQuantity: 2},
			expectedSecResult:   &UpdateProductOutput{RemainingQuantity: 0},
			expectedError:       nil,
		},
		{
			name:                "Given two product amount update when execute use case with amount greater than current stock should return an error",
			expectedStock:       product,
			firstReduce:         UpdateProductInput{ProductID: product.ID, Quantity: 3},
			secondReduce:        UpdateProductInput{ProductID: product.ID, Quantity: 7},
			expectedFirstResult: &UpdateProductOutput{RemainingQuantity: 2},
			expectedSecResult:   &UpdateProductOutput{RemainingQuantity: 2},
			expectedError:       entity.ErrInsufficientStock,
		},
	}

	for _, tt := range tests {
		db := mock_db.SetupDB()
		product.Quantity = 5
		t.Run(tt.name, func(t *testing.T) {
			stockGateway := database.NewStockDb(db)
			ctx := context.Background()
			uow := uow.NewUow(ctx, db)
			uow.Register("StockDb", func(tx *sql.Tx) interface{} {
				return database.NewStockDb(db)
			})
			uow.Register("ProductBalanceDB", func(tx *sql.Tx) interface{} {
				return database.NewProductBalanceDB(db)
			})

			err := stockGateway.Save(&product)
			assert.Nil(t, err)

			uc := NewProductBalanceUpdateUseCase(uow)
			output, err := uc.Execute(tt.firstReduce)
			assert.Equal(t, tt.expectedFirstResult, output)
			if tt.secondReduce == (UpdateProductInput{}) {
				assert.Equal(t, tt.expectedError, err)
			} else {
				output, err := uc.Execute(tt.secondReduce)
				assert.Equal(t, tt.expectedSecResult, output)
				assert.Equal(t, tt.expectedError, err)
			}

		})

		mock_db.TearDown(db)
	}
}
