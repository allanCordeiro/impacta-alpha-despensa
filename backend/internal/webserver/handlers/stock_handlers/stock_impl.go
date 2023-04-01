package stock_handlers

import "github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/gateway"

type StockHandler struct {
	StockGateway gateway.StockGateway
}

func NewStockandler(db gateway.StockGateway) *StockHandler {
	return &StockHandler{StockGateway: db}
}
