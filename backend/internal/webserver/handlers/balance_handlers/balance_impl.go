package balance_handlers

import (
	"github.com/AllanCordeiro/impacta-alpha-despensa/pkg/uow"
)

type BalanceHandler struct {
	ProductBalanceUow uow.UowInterface
}

func NewProductBalance(uow uow.UowInterface) *BalanceHandler {
	return &BalanceHandler{ProductBalanceUow: uow}
}
