package balance_handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type Input struct {
	Quantity int `json:"quantity"`
}

type Response struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

// CreateProductBalance godoc
// @Summary 			Create Product Balance
// @Description 		Create product stock transactions
// @Tags 				products
// @Accept 				json
// @Produce 			json
// @Param 				request body	Input	true	"product decrease amount"
// @Param 				productID path string true "product identifier"
// @Success 			200	{object}	Response
// @Failure 			400	{object}	Response
// @Failure 			500	{object}	Response
// @Router 				/api/products/{productID}/decrease [put]
func (h *BalanceHandler) CreateProductBalance(w http.ResponseWriter, r *http.Request) {
	var input usecase.UpdateProductInput
	var outResponse Response
	var req Input

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		if err.Error() != "EOF" {
			log.Println(err)
			outResponse.Status = "error"
			outResponse.StatusCode = http.StatusBadRequest
			outResponse.Data = "unrecognized request format"
			_ = json.NewEncoder(w).Encode(&outResponse)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	input.ProductID = chi.URLParam(r, "productID")
	input.Quantity = req.Quantity
	if input.Quantity == 0 {
		input.Quantity = 1
	}

	uc := usecase.NewProductBalanceUpdateUseCase(h.ProductBalanceUow)
	ouput, err := uc.Execute(input)
	if err != nil {
		log.Println(err)
		outResponse.Status = "error"
		if err == usecase.ErrInternal {
			outResponse.StatusCode = http.StatusInternalServerError
			outResponse.Data = "internal server error. please reach out support team"
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(&outResponse)
			return
		}
		if err.Error() == "product not found" {
			outResponse.StatusCode = http.StatusNotFound
			outResponse.Data = err.Error()
			w.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(w).Encode(&outResponse)
			return
		}
		outResponse.StatusCode = http.StatusBadRequest
		outResponse.Data = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(&outResponse)
		return
	}
	outResponse.Status = "success"
	outResponse.StatusCode = http.StatusOK
	outResponse.Data = &ouput
	_ = json.NewEncoder(w).Encode(&outResponse)
}
