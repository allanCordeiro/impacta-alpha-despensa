package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
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
	tests := []struct {
		testName     string
		product      usecase.CreateProductInput
		shouldBeOk   bool
		responseCode int
		responseData Response
	}{
		{
			testName: "Given a valid data when send to post handler should return ok",
			product: usecase.CreateProductInput{
				Name:           "Valid product",
				CreationDate:   time.Now().Format("2006-01-02"),
				Quantity:       2,
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
			product: usecase.CreateProductInput{
				Name:           "Invalid product",
				CreationDate:   time.Now().Add(time.Hour * 24 * 5).Format("2006-01-02"),
				Quantity:       2,
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
			product: usecase.CreateProductInput{
				Name:           "Invalid product",
				CreationDate:   time.Now().Format("2006-01-02"),
				Quantity:       2,
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
			product: usecase.CreateProductInput{
				Name:           "Invalid product",
				CreationDate:   time.Now().Format("2006-01-02"),
				Quantity:       0,
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
			product: usecase.CreateProductInput{
				Name:           "Invalid product",
				CreationDate:   time.Now().Add(time.Hour * 24 * 2).Format("2006-01-02"),
				Quantity:       2,
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
			fmt.Println(received.Data)
		})
	}

}
