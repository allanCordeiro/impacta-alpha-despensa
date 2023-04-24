package usecase

import (
	"log"
	"strconv"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/gateway"
)

type ProductOutput struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetStatisticsOutput struct {
	MinimalQuantity  int             `json:"minimal_quantity"`
	AffectedProducts int             `json:"affected_products"`
	ProductList      []ProductOutput `json:"product_list"`
}

type ProductStatisticsGetUseCase struct {
	StockGateway    gateway.StockGateway
	MinimalQuantity int
}

func NewProductStatisticsGetUseCase(stockGateway gateway.StockGateway, minimalQuantity int) *ProductStatisticsGetUseCase {
	return &ProductStatisticsGetUseCase{
		StockGateway:    stockGateway,
		MinimalQuantity: minimalQuantity,
	}
}

func (p *ProductStatisticsGetUseCase) Execute() GetStatisticsOutput {
	var output GetStatisticsOutput
	var prdOutput []ProductOutput

	entities, err := p.StockGateway.GetAllProducts()
	if err != nil {
		log.Println(err)
		return GetStatisticsOutput{
			MinimalQuantity:  p.MinimalQuantity,
			AffectedProducts: 0,
			ProductList:      []ProductOutput{},
		}
	}

	validProducts := getValidProducts(entities)
	output.MinimalQuantity = p.MinimalQuantity
	for _, prd := range validProducts {
		prdQuantity, _ := strconv.Atoi(prd.Quantity)
		if prdQuantity <= p.MinimalQuantity {
			prdOutput = append(prdOutput, ProductOutput{
				Id:   prd.ID,
				Name: prd.Name,
			})
			output.AffectedProducts += 1
		}
	}
	output.ProductList = prdOutput
	return output

}

func getValidProducts(productList []entity.Product) []GetProductOutput {
	var products []GetProductOutput
	for _, prd := range productList {
		var product GetProductOutput
		if isExpirationDateValid(prd.ExpirationDate) && hasEnoughQuantity(prd.Quantity) {
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
