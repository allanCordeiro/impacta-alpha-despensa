package balance_handlers

import (
	"encoding/json"
	"net/http"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/usecase"
)

type ProductListOutput struct {
	Name              string `json:"name"`
	DeductedAmount    int    `json:"deducted_quantity"`
	RemainingQuantity int    `json:"remaining_quantity"`
}

type ProductsBalanceOutput struct {
	Products map[string][]ProductListOutput
}

// CreateProductBalance godoc
// @Summary 			Get All Products Balance
// @Description 		Get all balance history
// @Tags 				products
// @Accept 				json
// @Produce 			json
// @Success 			200	{object}	Response
// @Failure 			404	{object}	Response
// @Failure 			500	{object}	Response
// @Router 				/api/products/balance [get]
func (h *BalanceHandlerWithGateway) GetBalance(w http.ResponseWriter, r *http.Request) {
	uc := usecase.NewProductBalanceUseCase(h.ProductBalanceGateway, h.StockGateway)
	output := uc.Execute()
	if len(output.Products) == 0 {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string][]ProductListOutput{})
		return
	}

	listOutput := make(map[string][]ProductListOutput)
	for key, productsBalance := range output.Products {
		var prdList []ProductListOutput
		for _, product := range productsBalance {
			prdList = append(prdList, ProductListOutput{
				Name:              product.Name,
				DeductedAmount:    product.DeductedAmount,
				RemainingQuantity: product.RemainingQuantity,
			})
		}
		listOutput[key] = append(listOutput[key], prdList...)
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&listOutput)
}
