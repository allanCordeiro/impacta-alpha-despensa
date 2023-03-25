package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEntity(t *testing.T) {
	scenarios := []struct {
		name           string
		product        string
		creationDate   time.Time
		quantity       int
		expirationDate time.Time
		isShouldBeOk   bool
		expectedErr    []error
	}{
		{
			name:           "given valid data when call NewProduct should be a valid result",
			product:        "Product 1",
			creationDate:   time.Now(),
			quantity:       16,
			expirationDate: time.Now().Add(time.Hour * 24 * 10),
			isShouldBeOk:   true,
			expectedErr:    nil,
		},
		{
			name:           "given future creation date when call NewProduct should send an error",
			product:        "Product 1",
			creationDate:   time.Now().Add(time.Hour * 24 * 3),
			quantity:       16,
			expirationDate: time.Now().Add(time.Hour * 24 * 10),
			isShouldBeOk:   false,
			expectedErr:    []error{ErrCreationDateInTheFuture},
		},
		{
			name:           "given expiration date in the past when call NewProduct should send an error",
			product:        "Product 1",
			creationDate:   time.Now(),
			quantity:       16,
			expirationDate: time.Now().Add(-time.Hour * 24 * 10),
			isShouldBeOk:   false,
			expectedErr:    []error{ErrExpirationDateInThePast},
		},
		{
			name:           "given negative quantity when call NewProduct should send an error",
			product:        "Product 1",
			creationDate:   time.Now(),
			quantity:       -5,
			expirationDate: time.Now().Add(time.Hour * 24 * 10),
			isShouldBeOk:   false,
			expectedErr:    []error{ErrInvalidQuantity},
		},
		{
			name:           "given greater than 33k quantity when call NewProduct should send an error",
			product:        "Product 1",
			creationDate:   time.Now(),
			quantity:       33000,
			expirationDate: time.Now().Add(time.Hour * 24 * 10),
			isShouldBeOk:   false,
			expectedErr:    []error{ErrInvalidQuantity},
		},
	}

	for _, tt := range scenarios {
		t.Run(tt.name, func(t *testing.T) {
			prd := NewProduct(tt.product, tt.creationDate, tt.quantity, tt.expirationDate)
			ok, err := prd.IsValid()
			assert.NotNil(t, prd.ID)
			assert.Equal(t, tt.expectedErr, err)
			assert.Equal(t, tt.isShouldBeOk, ok)
		})
	}
}
