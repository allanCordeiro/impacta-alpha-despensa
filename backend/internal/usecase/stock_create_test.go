package usecase

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProductUseCase_Execute(t *testing.T) {
	m := &mocks.StockGatewayMock{}
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
		{
			name:              "given quantity greater than 32600 when execute CreateProductUseCase should return an error",
			product:           "product 1",
			creationDate:      time.Now().Format("2006-01-02"),
			quantity:          35000,
			expirationDate:    time.Now().Add(time.Hour * 24 * 5).Format("2006-01-02"),
			isShouldBeOk:      true,
			expectedErrNumber: 1,
			expectedErr:       entity.ErrInvalidQuantity,
		},
		{
			name:              "given incorrect creation date when execute CreateProductUseCase should return an error",
			product:           "product 1",
			creationDate:      "20/07/2012",
			quantity:          -15,
			expirationDate:    time.Now().Add(time.Hour * 24 * 5).Format("2006-01-02"),
			isShouldBeOk:      true,
			expectedErrNumber: 1,
			expectedErr:       entity.ErrInvalidCreationDate,
		},
		{
			name:              "given quantity greater than 32600 when execute CreateProductUseCase should return an error",
			product:           "product 1",
			creationDate:      time.Now().Format("2006-01-02"),
			quantity:          35000,
			expirationDate:    "20/07/2077",
			isShouldBeOk:      true,
			expectedErrNumber: 1,
			expectedErr:       entity.ErrInvalidExpirationDate,
		},
	}

	for _, tt := range brScenarios {
		t.Run(tt.name, func(t *testing.T) {
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
				assert.Equal(t, tt.expectedErr.Error(), errMsg.Err)
			}
			m.AssertExpectations(t)
			m.AssertNumberOfCalls(t, "Save", 1)
		})
	}

	t.Run("given a valid data when execute CreateProductUseCase but an gateway error comes in should return an internal error", func(t *testing.T) {
		m := &mocks.StockGatewayMock{}
		m.On("Save", mock.Anything).Return(errors.New("sql: database is closed"))
		uc := NewCreateProductUseCase(m)
		input := CreateProductInput{
			Name:           "product 1",
			CreationDate:   time.Now().Format("2006-01-02"),
			Quantity:       2,
			ExpirationDate: time.Now().Add(time.Hour * 24 * 5).Format("2006-01-02"),
		}
		output := uc.Execute(input)
		fmt.Println(output)
		assert.NotNil(t, output.Msgs)
		for _, errMsg := range output.Msgs {
			assert.Equal(t, "internal", errMsg.Entity)
			assert.Equal(t, "sql: database is closed", errMsg.Err)
		}
		m.AssertExpectations(t)
		m.AssertNumberOfCalls(t, "Save", 1)
	})
}
