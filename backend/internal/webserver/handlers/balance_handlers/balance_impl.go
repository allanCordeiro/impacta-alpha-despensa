package balance_handlers

import (
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/gateway"
	"github.com/AllanCordeiro/impacta-alpha-despensa/pkg/uow"
)

type BalanceHandler struct {
	ProductBalanceUow uow.UowInterface
}

func NewProductBalance(uow uow.UowInterface) *BalanceHandler {
	return &BalanceHandler{ProductBalanceUow: uow}
}

type BalanceHandlerWithGateway struct {
	ProductBalanceGateway gateway.ProductBalanceGateway
	StockGateway          gateway.StockGateway
}

func NewProductBalanceWithGateway(balanceGateway gateway.ProductBalanceGateway, stockGateway gateway.StockGateway) *BalanceHandlerWithGateway {
	return &BalanceHandlerWithGateway{
		ProductBalanceGateway: balanceGateway,
		StockGateway:          stockGateway,
	}
}
