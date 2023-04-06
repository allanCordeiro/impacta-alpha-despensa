package stock_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/usecase"
)

// GetProducts godoc
// @Summary 			Get Products list
// @Description 		Get the list of all available products
// @Tags 				stock
// @Produce 			json
// @Success 			200	{object}	[]usecase.GetProductOutput
// @Failure 			500	{object}	Response
// @Router 				/api/stock [get]
func (h *StockHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	var output []usecase.GetProductOutput
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	uc := usecase.NewGetProductUseCase(h.StockGateway)
	output = uc.Execute()
	_ = json.NewEncoder(w).Encode(output)
}
