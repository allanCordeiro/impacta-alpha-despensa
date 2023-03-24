package usecase

import (
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type StockGetGatewayMock struct {
	mock.Mock
}

func (m *StockGetGatewayMock) Save(stock *entity.Product) error {
	return nil
}

func (m *StockGetGatewayMock) GetProducts() []entity.Product {
	return []entity.Product{}
}

func (m *StockGetGatewayMock) GetByID(id string) (*entity.Product, error) {
	return &entity.Product{}, nil
}

func (m *StockGetGatewayMock) GetAllProducts() ([]entity.Product, error) {
	args := m.Called()
	return args.Get(0).([]entity.Product), args.Error(1)
}

func TestGetProductUseCase_Execute(t *testing.T) {
	m := &StockGetGatewayMock{}
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
