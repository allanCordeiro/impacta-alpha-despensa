package usecase

import (
	"testing"
	"time"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetProductUseCase_Execute(t *testing.T) {
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
		{
			ID:             "3",
			Name:           "produto vencido",
			CreationDate:   time.Now(),
			ExpirationDate: time.Now().Add(-time.Hour * 5),
			Quantity:       2,
		},
		{
			ID:             "4",
			Name:           "produto com quantidade igual a zero",
			CreationDate:   time.Now(),
			ExpirationDate: time.Now().Add(time.Hour * 24 * 5),
			Quantity:       0,
		},
	}

	t.Run("Given a request to get products when customer request data should return available data", func(t *testing.T) {
		m.On("GetAllProducts").Return(productList, nil)
		uc := NewGetProductUseCase(m)
		output := uc.Execute()
		assert.Equal(t, 2, len(output))
		m.AssertExpectations(t)
		m.AssertNumberOfCalls(t, "GetAllProducts", 1)
	})
}
