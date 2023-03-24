package usecase

import (
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/gateway"
	"log"
	"strconv"
	"time"
)

type GetProductInput struct {
}

type GetProductOutput struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	CreationDate   string `json:"creation_date"`
	Quantity       string `json:"quantity"`
	ExpirationDate string `json:"expiration_date"`
}

type GetProductUseCase struct {
	StockGateway gateway.StockGateway
}

func NewGetProductUseCase(stockGateway gateway.StockGateway) *GetProductUseCase {
	return &GetProductUseCase{
		StockGateway: stockGateway,
	}
}

func (p *GetProductUseCase) Execute() []GetProductOutput {
	entities, err := p.StockGateway.GetAllProducts()
	if err != nil {
		log.Println(err)
		//todo better way to send an error msg to front
	}

	var products []GetProductOutput
	for _, prd := range entities {
		var product GetProductOutput
		if isExpirationDateValid(prd.ExpirationDate) {
			product = GetProductOutput{
				ID:             prd.ID,
				Name:           prd.Name,
				CreationDate:   prd.CreationDate.Format("2006-01-02"),
				Quantity:       strconv.Itoa(prd.Quantity),
				ExpirationDate: prd.ExpirationDate.Format("2006-01-02"),
			}
			products = append(products, product)
		}
	}
	return products
}

func isExpirationDateValid(date time.Time) bool {
	return date.After(time.Now())
}
