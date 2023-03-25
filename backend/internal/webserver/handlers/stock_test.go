package handlers

import (
	"encoding/json"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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
	return args.Get(0).([]entity.Product), nil
}

func TestStockHandler_GetProducts(t *testing.T) {
	listProducts := []entity.Product{
		{
			ID:             "1",
			Name:           "produto 1",
			CreationDate:   time.Now(),
			Quantity:       5,
			ExpirationDate: time.Now().Add(time.Hour * 24 * 5),
		},
		{
			ID:             "2",
			Name:           "produto 2",
			CreationDate:   time.Now(),
			Quantity:       5,
			ExpirationDate: time.Now().Add(time.Hour * 24 * 5),
		},
		{
			ID:             "3",
			Name:           "produto vencido",
			CreationDate:   time.Now(),
			Quantity:       5,
			ExpirationDate: time.Now().Add(-time.Hour * 24 * 5),
		},
	}

	t.Run("Given a list of products when call get method should display non-expired products list", func(t *testing.T) {
		m := &StockGetGatewayMock{}
		var received []usecase.GetProductOutput
		m.On("GetAllProducts").Return(listProducts)
		getHandler := NewStockandler(m)

		request, err := http.NewRequest(http.MethodGet, "/api/stock", nil)
		response := httptest.NewRecorder()
		assert.Nil(t, err)
		getHandler.GetProducts(response, request)
		err = json.Unmarshal(response.Body.Bytes(), &received)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, 2, len(received))

	})
}

func TestStockHandler_CreateProduct(t *testing.T) {

}
