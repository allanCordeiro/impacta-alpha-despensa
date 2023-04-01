package mocks

import (
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type StockGatewayMock struct {
	mock.Mock
}

func (m *StockGatewayMock) Save(stock *entity.Product) error {
	args := m.Called(stock)
	return args.Error(0)
}

func (m *StockGatewayMock) GetByID(id string) (*entity.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Product), args.Error(1)
}

func (m *StockGatewayMock) GetAllProducts() ([]entity.Product, error) {
	args := m.Called()
	return args.Get(0).([]entity.Product), args.Error(1)
}

func (m *StockGatewayMock) UpdateQuantity(stock *entity.Product) error {
	args := m.Called(stock)
	return args.Error(0)
}
