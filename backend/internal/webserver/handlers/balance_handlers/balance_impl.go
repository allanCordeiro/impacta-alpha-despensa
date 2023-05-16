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
}

func NewProductBalanceWithGateway(balanceGateway gateway.ProductBalanceGateway) *BalanceHandlerWithGateway {
	return &BalanceHandlerWithGateway{
		ProductBalanceGateway: balanceGateway,
	}
}
