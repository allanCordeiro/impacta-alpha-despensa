package usecase

import (
	"testing"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductBalanceUpdate(t *testing.T) {
	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	uc := NewProductBalanceUpdateUseCase(mockUow)
	_, err := uc.Execute(UpdateProductInput{ProductID: "id", Quantity: 2})
	assert.Nil(t, err)
}
