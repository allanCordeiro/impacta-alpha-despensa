package stock_handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/usecase"
)

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
// @Param 				request body	usecase.CreateProductInput	true	"product information"
// @Success 			200	{object}	Response
// @Failure 			400	{object}	Response
// @Failure 			500	{object}	Response
// @Router 				/api/stock [post]
func (h *StockHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateProductInput
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
	output = uc.Execute(input)
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
