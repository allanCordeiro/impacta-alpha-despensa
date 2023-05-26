package usecase

import (
	"testing"
	"time"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetProductBalance(t *testing.T) {
	stockMock := &mocks.StockGatewayMock{}
	balanceMock := &mocks.ProductBalance{}
	product := []entity.Product{
		{
			ID:             "prd 1",
			Name:           "produto 1",
			CreationDate:   time.Now(),
			ExpirationDate: time.Now().Add(time.Hour * 24 * 5),
			Quantity:       1,
		},
	}

	balance := []entity.ProductBalance{
		{
			ProductID:         "prd 1",
			DeductedAmount:    5,
			RemainingQuantity: 10,
			DeductedDate:      time.Now().Add(-time.Hour * 24 * 5),
		},
		{
			ProductID:         "prd 1",
			DeductedAmount:    3,
			RemainingQuantity: 7,
			DeductedDate:      time.Now().Add(-time.Hour * 24 * 5),
		},
		{
			ProductID:         "prd 1",
			DeductedAmount:    1,
			RemainingQuantity: 6,
			DeductedDate:      time.Now().Add(-time.Hour * 24 * 2),
		},
		{
			ProductID:         "prd 1",
			DeductedAmount:    1,
			RemainingQuantity: 6,
			DeductedDate:      time.Now().Add(-time.Hour * 24 * 1),
		},
	}

	t.Run("Given following data when call the use case should return 3 items grouped by date", func(t *testing.T) {
		stockMock.On("GetAllProducts").Return(product, nil)
		balanceMock.On("GetAllProductsBalance", mock.Anything).Return(balance, nil)
		expectedLenght := 3
		expectedWeekAgoLenght := 2
		exepectedWeekAgoDate := time.Now().Add(-time.Hour * 24 * 5).Format("02-01-2006")

		uc := NewProductBalanceUseCase(balanceMock, stockMock)
		output := uc.Execute()
		assert.NotNil(t, output)
		assert.Equal(t, expectedLenght, len(output.Products))
		assert.NotNil(t, output.Products[exepectedWeekAgoDate])
		assert.Equal(t, expectedWeekAgoLenght, len(output.Products[exepectedWeekAgoDate]))
	})
}
