package handlers

import (
	"encoding/json"
	"errors"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/gateway"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/usecase"
	"log"
	"net/http"
)

var (
	ErrInvalidContent = errors.New("conteudo invalido")
)

type Response struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

type StockHandler struct {
	StockGateway gateway.StockGateway
}

func NewStockandler(db gateway.StockGateway) *StockHandler {
	return &StockHandler{StockGateway: db}
}

func (h *StockHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateProductInput
	var output usecase.CreateProductOutput

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&input)
	log.Println(err)
	log.Println(input)
	if err != nil {
		output.Msgs = append(output.Msgs, usecase.Msg{Entity: "stock", Err: ErrInvalidContent.Error()})
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(&Response{
			Status:     "error",
			StatusCode: http.StatusBadRequest,
			Data:       &output,
		})
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	uc := usecase.NewCreateProductUseCase(h.StockGateway)
	out := uc.Execute(input)
	if len(out.Msgs) == 0 {
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(&Response{
			Status:     "success",
			StatusCode: http.StatusOK,
		})
		return
	}

	for _, errMsg := range out.Msgs {
		if errMsg.Entity == "internal" {
			log.Print(errMsg.Err)
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(&Response{
				Status:     "error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		err = json.NewEncoder(w).Encode(&output)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(&Response{
		Status:     "error",
		StatusCode: http.StatusInternalServerError,
		Data:       &out,
	})
	return
}
