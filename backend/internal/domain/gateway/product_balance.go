package gateway

import "github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"

type ProductBalanceGateway interface {
	Save(productBalance *entity.ProductBalance) error
	GetByProductId(productId string) ([]entity.ProductBalance, error)
}
