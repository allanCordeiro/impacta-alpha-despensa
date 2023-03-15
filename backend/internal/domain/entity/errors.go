package entity

import "errors"

var (
	ErrCreationDateInTheFuture = errors.New("a data de entrada nao deve estar no futuro")
	//ErrInvalidCreationDate     = errors.New("a data de entrada esta invalida")
	ErrExpirationDateInThePast = errors.New("data de vencimento anterior ao dia atual")
	//ErrInvalidExpirationDate   = errors.New("a data de vencimento esta invalida")
	ErrInvalidQuantity = errors.New("a quantidade de itens esta incorreta")
)
