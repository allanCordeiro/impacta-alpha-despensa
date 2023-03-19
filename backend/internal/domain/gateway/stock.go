package gateway

import "github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"

type StockGateway interface {
	Save(stock *entity.Product) error
	GetByID(id string) (*entity.Product, error)
	GetAllProducts() ([]entity.Product, error)
}
