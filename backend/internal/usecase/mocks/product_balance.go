package mocks

import (
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type ProductBalance struct {
	mock.Mock
}

func (m *ProductBalance) Save(productBalance *entity.ProductBalance) error {
	args := m.Called(productBalance)
	return args.Error(0)
}

func (m *ProductBalance) GetByProductId(productId string) ([]entity.ProductBalance, error) {
	args := m.Called(productId)
	return args.Get(0).([]entity.ProductBalance), args.Error(1)
}
