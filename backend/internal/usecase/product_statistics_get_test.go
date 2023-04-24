package usecase

import (
	"testing"
	"time"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetProductStatistics_Execute(t *testing.T) {
	m := &mocks.StockGatewayMock{}

	productList := []entity.Product{
		{
			ID:             "1",
			Name:           "produto 1",
			CreationDate:   time.Now(),
			ExpirationDate: time.Now().Add(time.Hour * 24 * 5),
			Quantity:       3,
		},
		{
			ID:             "2",
			Name:           "produto 2",
			CreationDate:   time.Now(),
			ExpirationDate: time.Now().Add(time.Hour * 24 * 5),
			Quantity:       2,
		},
	}

	t.Run("Given a call to statistics when products fullfill the requirements should be dsiplayed", func(t *testing.T) {
		expectedMinimalQuantity := 1
		expectedAffectedProducts := 2

		m.On("GetAllProducts").Return(productList, nil)
		uc := NewProductStatisticsGetUseCase(m, 1)

		output := uc.Execute()

		assert.NotNil(t, output)
		assert.Equal(t, expectedMinimalQuantity, output.MinimalQuantity)
		assert.Equal(t, expectedAffectedProducts, output.AffectedProducts)
		assert.Equal(t, expectedAffectedProducts, len(output.ProductList))
	})
}
