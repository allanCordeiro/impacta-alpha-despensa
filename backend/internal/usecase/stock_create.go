package usecase

import (
	"time"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/gateway"
)

type CreateProductInput struct {
	Name           string `json:"name"`
	CreationDate   string `json:"creation_date"`
	Quantity       int    `json:"quantity"`
	ExpirationDate string `json:"expiration_date"`
}

type Msg struct {
	Entity string `json:"entity"`
	Err    string `json:"err"`
}

type CreateProductOutput struct {
	Msgs []Msg `json:"errors"`
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
	if !isInputValid(input) {
		errorMsg := Msg{Entity: "stock", Err: entity.ErrInvalidStruct.Error()}
		errors.Msgs = append(errors.Msgs, errorMsg)
	}
	creationDate, err := dateParse(input.CreationDate)
	if err != nil {
		errorMsg := Msg{Entity: "stock", Err: entity.ErrInvalidCreationDate.Error()}
		errors.Msgs = append(errors.Msgs, errorMsg)
	}

	expirationDate, err := dateParse(input.ExpirationDate)
	if err != nil {
		errorMsg := Msg{Entity: "stock", Err: entity.ErrInvalidExpirationDate.Error()}
		errors.Msgs = append(errors.Msgs, errorMsg)
	}
	if !errors.shouldProceed() {
		return errors
	}

	prd := entity.NewProduct(input.Name, creationDate, input.Quantity, expirationDate)
	if ok, err := prd.IsValid(); !ok {
		for _, singleErr := range err {
			errorMsg := Msg{Entity: "stock", Err: singleErr.Error()}
			errors.Msgs = append(errors.Msgs, errorMsg)
		}
	}

	if !errors.shouldProceed() {
		return errors
	}

	err = p.StockGateway.Save(prd)
	if err != nil {
		errorMsg := Msg{Entity: "internal", Err: err.Error()}
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
	return len(e.Msgs) <= 0
}

func isInputValid(input CreateProductInput) bool {
	if len(input.Name) == 0 {
		return false
	}
	if input.CreationDate == "" || input.ExpirationDate == "" {
		return false
	}
	return true
}
