package balance_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type ProductBalanceInput struct {
	ProductID string
}

type OperationList struct {
	OperationDate     string `json:"operation_date"`
	DeductedQuantity  int    `json:"deducted_quantity"`
	RemainingQuantity int    `json:"reimaining_quantity"`
}

// CreateProductBalance godoc
// @Summary 			Get Product Balance
// @Description 		Get product history of reductions
// @Tags 				products
// @Accept 				json
// @Produce 			json
// @Param 				productID path string true "product identifier"
// @Success 			200	{object}	Response
// @Failure 			404	{object}	Response
// @Failure 			500	{object}	Response
// @Router 				/api/products/{productID}/balance [get]
func (h *BalanceHandlerWithGateway) GetProductBalance(w http.ResponseWriter, r *http.Request) {
	input := ProductBalanceInput{
		ProductID: chi.URLParam(r, "productID"),
	}

	uc := usecase.NewProductBalanceGetByIDUseCase(h.ProductBalanceGateway)
	output := uc.Execute(input.ProductID)
	if len(output.BalanceList) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var listOutput []OperationList
	for _, prd := range output.BalanceList {
		operation := OperationList{
			OperationDate:     prd.OperationDate,
			DeductedQuantity:  prd.DeductedQuantity,
			RemainingQuantity: prd.RemainingQuantity,
		}
		listOutput = append(listOutput, operation)
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&listOutput)
}
