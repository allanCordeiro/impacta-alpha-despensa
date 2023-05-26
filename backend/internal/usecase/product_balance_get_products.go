package usecase

import (
	"log"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/gateway"
)

type ProductList struct {
	Name              string
	DeductedAmount    int
	RemainingQuantity int
}

type ProductsBalanceOutput struct {
	Products map[string][]ProductList
}

type ProductBalanceUseCase struct {
	ProductBalanceGateway gateway.ProductBalanceGateway
	StockGateway          gateway.StockGateway
}

func NewProductBalanceUseCase(pb gateway.ProductBalanceGateway, stockGateway gateway.StockGateway) *ProductBalanceUseCase {
	return &ProductBalanceUseCase{
		ProductBalanceGateway: pb,
		StockGateway:          stockGateway,
	}
}

func (u *ProductBalanceUseCase) Execute() ProductsBalanceOutput {
	entities, err := u.ProductBalanceGateway.GetAllProductsBalance()
	if err != nil {
		log.Println(err)
		return ProductsBalanceOutput{}
	}
	prodByData := make(map[string][]ProductList)
	productList := u.getProductList()
	var keys []string

	for _, balance := range entities {
		date := balance.DeductedDate.Format("02-01-2006")

		productHistory := ProductList{
			Name:              productList[balance.ProductID],
			RemainingQuantity: balance.RemainingQuantity,
			DeductedAmount:    balance.DeductedAmount,
		}
		prodByData[date] = append(prodByData[date], productHistory)
		keys = append(keys, date)
	}

	// Reverter a ordem das chaves
	reversedKeys := make([]string, len(keys))
	for i, j := 0, len(keys)-1; i < len(keys); i, j = i+1, j-1 {
		reversedKeys[i] = keys[j]
	}

	output := make(map[string][]ProductList)
	for _, k := range reversedKeys {
		output[k] = prodByData[k]
	}

	return ProductsBalanceOutput{Products: output}
}

func (u *ProductBalanceUseCase) getProductList() map[string]string {
	productData := make(map[string]string)
	products, err := u.StockGateway.GetAllProducts()
	if err != nil {
		log.Panic(err)
	}
	for _, prd := range products {
		productData[prd.ID] = prd.Name
	}
	return productData
}
