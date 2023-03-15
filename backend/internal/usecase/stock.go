package usecase

import (
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

type Msg struct {
	Entity string
	Err    error
}

type CreateProductOutput struct {
	Msgs []Msg
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
	errors := CreateProductOutput{}
	creationDate, err := dateParse(input.CreationDate)
	if err != nil {
		errorMsg := Msg{Entity: "stock", Err: entity.ErrInvalidCreationDate}
		errors.Msgs = append(errors.Msgs, errorMsg)
	}

	expirationDate, err := dateParse(input.ExpirationDate)
	if err != nil {
		errorMsg := Msg{Entity: "stock", Err: entity.ErrInvalidExpirationDate}
		errors.Msgs = append(errors.Msgs, errorMsg)
	}
	if !errors.shouldProceed() {
		return errors
	}

	prd := entity.NewProduct(input.Name, creationDate, input.Quantity, expirationDate)
	if ok, err := prd.IsValid(); !ok {
		errorMsg := Msg{Entity: "stock", Err: err}
		errors.Msgs = append(errors.Msgs, errorMsg)
	}

	if !errors.shouldProceed() {
		return errors
	}

	err = p.StockGateway.Save(prd)
	if err != nil {
		errorMsg := Msg{Entity: "internal", Err: err}
		errors.Msgs = append(errors.Msgs, errorMsg)
	}
	return errors
}

func dateParse(date string) (time.Time, error) {
	rightDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}, err
	}
	return rightDate, nil
}

func (e *CreateProductOutput) shouldProceed() bool {
	if len(e.Msgs) > 0 {
		return false
	}
	return true
}
