package balance_handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/usecase"
	"github.com/go-chi/chi/v5"
)

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
// @Param 				request body	usecase.UpdateProductInput	true	"product decrease amount"
// @Success 			200	{object}	Response
// @Failure 			400	{object}	Response
// @Failure 			500	{object}	Response
// @Router 				/api/stock-decrease/{productID} [post]
func (h *BalanceHandler) CreateProductBalance(w http.ResponseWriter, r *http.Request) {
	var input usecase.UpdateProductInput

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	input.ProductID = chi.URLParam(r, "productID")
	if input.Quantity == 0 {
		input.Quantity = 1
	}

	uc := usecase.NewProductBalanceUpdateUseCase(h.ProductBalanceUow)
	ouput, err := uc.Execute(input)
	if err != nil {
		log.Println("erro de usecase")
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_ = json.NewEncoder(w).Encode(&ouput)
}
