package usecase

import (
	"fmt"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/gateway"
	"time"
)

type CreateProductInput struct {
	Name           string `json:"name"`
	CreationDate   string `json:"creation_date"`
	Quantity       int    `json:"quantity"`
	ExpirationDate string `json:"expiration_date"`
}

type CreateProductOutput struct {
	Msgs []Msg
}

type Msg struct {
	Entity         string
	ErrDescription string
}

type CreateProductUseCase struct {
	StockGateway gateway.StockGateway
}

func NewCreateProductUseCase(stockGateway gateway.StockGateway) *CreateProductUseCase {
	return &CreateProductUseCase{
		StockGateway: stockGateway,
	}
}

func (p *CreateProductUseCase) Execute(input CreateProductInput) CreateProductOutput {
	creationDate, err := dateParse(input.CreationDate)
	if err != nil {
		fmt.Println(err)
	}
	expirationDate, err := dateParse(input.ExpirationDate)
	if err != nil {
		fmt.Println(err)
	}
	prd := entity.NewProduct(input.Name, creationDate, input.Quantity, expirationDate)
	if ok, err := prd.IsValid(); ok {
		fmt.Println(err)
	}
	err = p.StockGateway.Save(prd)
	if err != nil {
		fmt.Println(err)
	}
	return CreateProductOutput{}
}

func dateParse(date string) (time.Time, error) {
	rightDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, err
	}
	return rightDate, nil
}
