package stock_handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func TestStockStatisticsHandler_GetStatistics(t *testing.T) {
	m := &StockGetGatewayMock{}
	productList := []entity.Product{
		{
			ID:             "1",
			Name:           "produto 1",
			CreationDate:   time.Now(),
			ExpirationDate: time.Now().Add(time.Hour * 24 * 5),
			Quantity:       1,
		},
		{
			ID:             "2",
			Name:           "produto 2",
			CreationDate:   time.Now(),
			ExpirationDate: time.Now().Add(time.Hour * 24 * 5),
			Quantity:       1,
		},
		{
			ID:             "3",
			Name:           "produto 3",
			CreationDate:   time.Now(),
			ExpirationDate: time.Now().Add(time.Hour * 24 * 5),
			Quantity:       5,
		},
	}

	t.Run("Given a call to statistics, when retrieve data should shown only products with minimal quantity", func(t *testing.T) {
		expectedMinQuantity := 1
		expectedAffectedProducts := 2
		var output usecase.GetStatisticsOutput
		m.On("GetAllProducts").Return(productList, nil)
		statisticsHander := NewStockandler(m)

		ctx := context.Background()
		ctx = context.WithValue(ctx, "min-quantity", 1)

		request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/api/stock/statistics", nil)

		response := httptest.NewRecorder()
		assert.Nil(t, err)

		statisticsHander.GetStatistics(response, request)
		err = json.Unmarshal(response.Body.Bytes(), &output)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expectedMinQuantity, output.MinimalQuantity)
		assert.Equal(t, expectedAffectedProducts, output.AffectedProducts)
		assert.Equal(t, expectedAffectedProducts, len(output.ProductList))
	})
}
