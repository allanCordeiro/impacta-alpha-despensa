package balance_handlers

import (
	"encoding/json"
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
// @Tags 				stock-decrease
// @Accept 				json
// @Produce 			json
// @Param 				request body	Input	true	"product decrease amount"
// @Success 			200	{object}	Response
// @Failure 			400	{object}	Response
// @Failure 			500	{object}	Response
// @Router 				/api/stock-decrease/{productID} [post]
func (h *BalanceHandler) CreateProductBalance(w http.ResponseWriter, r *http.Request) {
	var input usecase.UpdateProductInput
	var req Input

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		if err.Error() != "EOF" {
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
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_ = json.NewEncoder(w).Encode(&ouput)
}
