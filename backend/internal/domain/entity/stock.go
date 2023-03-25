package entity

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID             string
	Name           string
	CreationDate   time.Time
	Quantity       int
	ExpirationDate time.Time
}

func NewProduct(name string, creationDate time.Time, quantity int, expirationDate time.Time) *Product {
	return &Product{
		ID:             uuid.New().String(),
		Name:           name,
		CreationDate:   creationDate,
		Quantity:       quantity,
		ExpirationDate: expirationDate,
	}
}

func (p *Product) IsValid() (bool, []error) {
	var errorList []error

	if p.Quantity < 1 || p.Quantity > 32767 {
		errorList = append(errorList, ErrInvalidQuantity)
	}

	if p.CreationDate.After(time.Now()) {
		errorList = append(errorList, ErrCreationDateInTheFuture)
	}

	if p.ExpirationDate.Before(time.Now()) {
		errorList = append(errorList, ErrExpirationDateInThePast)
	}

	if len(errorList) > 0 {
		return false, errorList
	}

	return true, nil
}
