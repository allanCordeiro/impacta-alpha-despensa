package stock_handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/usecase"
)

type RequestInput struct {
	Name           string `json:"name"`
	CreationDate   string `json:"creation_date"`
	Quantity       string `json:"quantity"`
	ExpirationDate string `json:"expiration_date"`
}

type Response struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

// CreateProduct godoc
// @Summary 			Create Product
// @Description 		Create a new product and merge it to stock
// @Tags 				stock
// @Accept 				json
// @Produce 			json
// @Param 				request body	RequestInput	true	"product information"
// @Success 			200	{object}	Response
// @Failure 			400	{object}	Response
// @Failure 			500	{object}	Response
// @Router 				/api/stock [post]
func (h *StockHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var input RequestInput

	var output usecase.CreateProductOutput

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Response{
			Status:     "error",
			StatusCode: http.StatusBadRequest,
			Data:       nil,
		})
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	uc := usecase.NewCreateProductUseCase(h.StockGateway)
	qtyToInt, _ := strconv.Atoi(input.Quantity)
	quantityParsed := usecase.CreateProductInput{
		Name:           input.Name,
		CreationDate:   input.CreationDate,
		ExpirationDate: input.ExpirationDate,
		Quantity:       qtyToInt,
	}
	output = uc.Execute(quantityParsed)
	if len(output.Msgs) == 0 {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(&Response{
			Status:     "success",
			StatusCode: http.StatusOK,
		})
		return
	}

	for _, errMsg := range output.Msgs {
		if errMsg.Entity == "internal" {
			log.Print(errMsg.Err)
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(&Response{
				Status:     "error",
				StatusCode: http.StatusInternalServerError,
				Data:       "ocorreu um erro interno",
			})
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(&Response{
		Status:     "error",
		StatusCode: http.StatusBadRequest,
		Data:       &output,
	})
}
