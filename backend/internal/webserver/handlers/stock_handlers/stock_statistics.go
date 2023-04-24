package stock_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/usecase"
)

// GetProducts godoc
// @Summary 			Get Products statistics
// @Description 		Get statistics of how many products are below the minimal quantity in stock
// @Tags 				stock
// @Produce 			json
// @Success 			200	{object}	usecase.GetStatisticsOutput
// @Failure 			500	{object}	Response
// @Router 				/api/stock/statistics [get]
func (h *StockHandler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	var output usecase.GetStatisticsOutput
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	uc := usecase.NewProductStatisticsGetUseCase(h.StockGateway, 1)
	output = uc.Execute()

	_ = json.NewEncoder(w).Encode(output)
}
