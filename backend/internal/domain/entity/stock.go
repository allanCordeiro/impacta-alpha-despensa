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

func (p *Product) IsValid() (bool, error) {
	if p.Quantity < 1 || p.Quantity > 32767 {
		return false, ErrInvalidQuantity
	}

	if p.CreationDate.After(time.Now()) {
		return false, ErrCreationDateInTheFuture
	}

	if p.ExpirationDate.Before(time.Now()) {
		return false, ErrExpirationDateInThePast
	}

	return true, nil
}
