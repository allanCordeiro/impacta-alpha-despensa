package stock_handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/usecase"
	"github.com/stretchr/testify/assert"
)

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
		m.On("GetAllProducts").Return(listProducts, nil)
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
