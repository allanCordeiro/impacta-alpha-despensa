package gateway

import "github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"

type StockGateway interface {
	Save(stock entity.Product) error
	GetProducts() []entity.Product
	GetByID(id string) (*entity.Product, error)
}
