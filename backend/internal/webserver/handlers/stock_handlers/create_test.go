package stock_handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStockHandler_CreateProduct(t *testing.T) {
	tests := []struct {
		testName     string
		product      RequestInput
		shouldBeOk   bool
		responseCode int
		responseData Response
	}{
		{
			testName: "Given a valid data when send to post handler should return ok",
			product: RequestInput{
				Name:           "Valid product",
				CreationDate:   time.Now().Format("2006-01-02"),
				Quantity:       "2",
				ExpirationDate: time.Now().Add(time.Hour * 24 * 5).Format("2006-01-02"),
			},
			shouldBeOk:   true,
			responseCode: 200,
			responseData: Response{
				Status:     "success",
				StatusCode: http.StatusOK,
				Data:       nil,
			},
		},
		{
			testName: "Given a future creation date product when send to post handler should return error",
			product: RequestInput{
				Name:           "Invalid product",
				CreationDate:   time.Now().Add(time.Hour * 24 * 5).Format("2006-01-02"),
				Quantity:       "2",
				ExpirationDate: time.Now().Add(time.Hour * 24 * 5).Format("2006-01-02"),
			},
			shouldBeOk:   false,
			responseCode: 400,
			responseData: Response{
				Status:     "error",
				StatusCode: http.StatusBadRequest,
				Data: map[string]interface{}{
					"errors": []interface{}{
						map[string]interface{}{
							"entity": "stock",
							"err":    "a data de entrada nao deve estar no futuro",
						},
					},
				},
			},
		},
		{
			testName: "Given a expired product when send to post handler should return error",
			product: RequestInput{
				Name:           "Invalid product",
				CreationDate:   time.Now().Format("2006-01-02"),
				Quantity:       "2",
				ExpirationDate: time.Now().Add(-time.Hour * 24 * 5).Format("2006-01-02"),
			},
			shouldBeOk:   false,
			responseCode: 400,
			responseData: Response{
				Status:     "error",
				StatusCode: http.StatusBadRequest,
				Data: map[string]interface{}{
					"errors": []interface{}{
						map[string]interface{}{
							"entity": "stock",
							"err":    "data de vencimento anterior ao dia atual",
						},
					},
				},
			},
		},
		{
			testName: "Given a wrong product quantity when send to post handler should return error",
			product: RequestInput{
				Name:           "Invalid product",
				CreationDate:   time.Now().Format("2006-01-02"),
				Quantity:       "0",
				ExpirationDate: time.Now().Add(time.Hour * 24 * 5).Format("2006-01-02"),
			},
			shouldBeOk:   false,
			responseCode: 400,
			responseData: Response{
				Status:     "error",
				StatusCode: http.StatusBadRequest,
				Data: map[string]interface{}{
					"errors": []interface{}{
						map[string]interface{}{
							"entity": "stock",
							"err":    "a quantidade de itens esta incorreta",
						},
					},
				},
			},
		},
		{
			testName: "Given a product with several errors when send to post handler should return error's list",
			product: RequestInput{
				Name:           "Invalid product",
				CreationDate:   time.Now().Add(time.Hour * 24 * 2).Format("2006-01-02"),
				Quantity:       "0",
				ExpirationDate: time.Now().Add(-time.Hour * 24 * 5).Format("2006-01-02"),
			},
			shouldBeOk:   false,
			responseCode: 400,
			responseData: Response{
				Status:     "error",
				StatusCode: http.StatusBadRequest,
				Data: map[string]interface{}{
					"errors": []interface{}{
						map[string]interface{}{
							"entity": "stock",
							"err":    "a quantidade de itens esta incorreta",
						},
						map[string]interface{}{
							"entity": "stock",
							"err":    "a data de entrada nao deve estar no futuro",
						},
						map[string]interface{}{
							"entity": "stock",
							"err":    "data de vencimento anterior ao dia atual",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			m := &StockGetGatewayMock{}
			m.On("Save", mock.Anything).Return(nil)
			postHandler := NewStockandler(m)
			var received Response

			var buf bytes.Buffer
			json.NewEncoder(&buf).Encode(tt.product)
			request, err := http.NewRequest(http.MethodPost, "/api/stock", &buf)
			response := httptest.NewRecorder()
			assert.Nil(t, err)

			postHandler.CreateProduct(response, request)
			assert.Equal(t, tt.responseCode, response.Code)
			body, err := io.ReadAll(response.Body)
			assert.Nil(t, err)

			err = json.Unmarshal(body, &received)
			assert.Nil(t, err)
			assert.Equal(t, tt.responseData.Status, received.Status)
			assert.Equal(t, tt.responseData.StatusCode, received.StatusCode)
			assert.Equal(t, tt.responseData.Data, received.Data)
		})
	}

	t.Run("Given a valid product, when return any unexpected error should return an http error 500", func(t *testing.T) {
		m := &StockGetGatewayMock{}
		m.On("Save", mock.Anything).Return(errors.New("SQL whatever error"))
		product := RequestInput{
			Name:           "Valid product",
			CreationDate:   time.Now().Format("2006-01-02"),
			Quantity:       "2",
			ExpirationDate: time.Now().Add(time.Hour * 24 * 5).Format("2006-01-02"),
		}
		expectedResponse := "ocorreu um erro interno"

		postHandler := NewStockandler(m)
		var received Response

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(product)
		request, err := http.NewRequest(http.MethodPost, "/api/stock", &buf)
		response := httptest.NewRecorder()
		assert.Nil(t, err)

		postHandler.CreateProduct(response, request)
		body, _ := io.ReadAll(response.Body)
		err = json.Unmarshal(body, &received)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, "error", received.Status)
		assert.Equal(t, http.StatusInternalServerError, received.StatusCode)
		assert.Equal(t, expectedResponse, received.Data)

	})

}
