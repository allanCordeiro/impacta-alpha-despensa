package stock_handlers

import (
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

// Mocks
type StockGetGatewayMock struct {
	mock.Mock
}

func (m *StockGetGatewayMock) Save(stock *entity.Product) error {
	args := m.Called(stock)
	return args.Error(0)
}

func (m *StockGetGatewayMock) GetByID(id string) (*entity.Product, error) {
	return &entity.Product{}, nil
}

func (m *StockGetGatewayMock) GetAllProducts() ([]entity.Product, error) {
	args := m.Called()
	return args.Get(0).([]entity.Product), args.Error(1)
}

func (m *StockGetGatewayMock) UpdateQuantity(stock *entity.Product) error {
	args := m.Called(stock)
	return args.Error(0)
}
