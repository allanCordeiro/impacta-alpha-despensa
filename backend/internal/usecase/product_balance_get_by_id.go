package usecase

import (
	"log"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/gateway"
)

type ProductInput struct {
	ProductID string `json:"id"`
}

type OperationList struct {
	OperationDate     string `json:"operation_date"`
	DeductedQuantity  int    `json:"deducted_quantity"`
	RemainingQuantity int    `json:"reimaining_quantity"`
}

type ProductBalanceOutput struct {
	BalanceList []OperationList `json:"balance_list"`
}

type ProductBalanceGetByIDUseCase struct {
	ProductBalanceGateway gateway.ProductBalanceGateway
}

func NewProductBalanceGetByIDUseCase(productBalanceGateway gateway.ProductBalanceGateway) *ProductBalanceGetByIDUseCase {
	return &ProductBalanceGetByIDUseCase{
		ProductBalanceGateway: productBalanceGateway,
	}
}

func (u *ProductBalanceGetByIDUseCase) Execute(productID string) *ProductBalanceOutput {
	entities, err := u.ProductBalanceGateway.GetByProductId(productID)
	if err != nil {
		log.Println(err)
		return &ProductBalanceOutput{
			BalanceList: []OperationList{},
		}
	}

	var balance []OperationList

	for _, list := range entities {
		var balanceItem OperationList
		balanceItem.DeductedQuantity = list.DeductedAmount
		balanceItem.RemainingQuantity = list.RemainingQuantity
		balanceItem.OperationDate = list.DeductedDate.Format("2006-01-02 15:04:05")

		balance = append(balance, balanceItem)
	}
	return &ProductBalanceOutput{
		BalanceList: balance,
	}
}
