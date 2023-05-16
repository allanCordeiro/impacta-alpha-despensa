package usecase

import (
	"testing"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductBalanceByID(t *testing.T) {
	prd1bal1 := entity.NewProductBalance("prod 1", 1, 9)
	prd1bal2 := entity.NewProductBalance("prod 1", 2, 7)
	productBalance := []entity.ProductBalance{
		*prd1bal1,
		*prd1bal2,
	}

	t.Run("Given 3 prd balances with 2 distincts products when call balance by id should return only its respective data", func(t *testing.T) {
		m := &mocks.ProductBalance{}
		m.On("GetByProductId", mock.Anything).Return(productBalance, nil)

		uc := NewProductBalanceGetByIDUseCase(m)
		output := uc.Execute("prod 1")
		assert.NotNil(t, output)
		assert.Equal(t, 2, len(output.BalanceList))
	})
}
