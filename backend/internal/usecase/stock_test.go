package usecase

import (
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type StockGatewayMock struct {
	mock.Mock
}

func (m *StockGatewayMock) Save(stock *entity.Product) error {
	args := m.Called(stock)
	return args.Error(0)
}

func (m *StockGatewayMock) GetProducts() []entity.Product {
	return []entity.Product{}
}

func (m *StockGatewayMock) GetByID(id string) (*entity.Product, error) {
	return &entity.Product{}, nil
}

func TestCreateProductUseCase_Execute(t *testing.T) {
	m := &StockGatewayMock{}
	brScenarios := []struct {
		name              string
		product           string
		creationDate      string
		quantity          int
		expirationDate    string
		isShouldBeOk      bool
		expectedErrNumber int
		expectedErr       error
	}{
		{
			name:              "given correct data when execute CreateProductUseCase should return ok",
			product:           "product 1",
			creationDate:      time.Now().Format("2006-01-02"),
			quantity:          15,
			expirationDate:    time.Now().Add(time.Hour * 24 * 5).Format("2006-01-02"),
			isShouldBeOk:      true,
			expectedErrNumber: 0,
			expectedErr:       nil,
		},
		{
			name:              "given incorrect quantity when execute CreateProductUseCase should return an error",
			product:           "product 1",
			creationDate:      time.Now().Format("2006-01-02"),
			quantity:          -15,
			expirationDate:    time.Now().Add(time.Hour * 24 * 5).Format("2006-01-02"),
			isShouldBeOk:      true,
			expectedErrNumber: 1,
			expectedErr:       entity.ErrInvalidQuantity,
		},
	}

	for _, tt := range brScenarios {
		m.On("Save", mock.Anything).Return(nil)
		uc := NewCreateProductUseCase(m)
		input := CreateProductInput{
			Name:           tt.product,
			CreationDate:   tt.creationDate,
			Quantity:       tt.quantity,
			ExpirationDate: tt.expirationDate,
		}
		output := uc.Execute(input)
		assert.Equal(t, tt.expectedErrNumber, len(output.Msgs))
		for _, errMsg := range output.Msgs {
			assert.Equal(t, tt.expectedErr, errMsg.Err)
		}
	}
}
