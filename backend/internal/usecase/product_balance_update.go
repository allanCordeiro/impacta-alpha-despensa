package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/gateway"
	"github.com/AllanCordeiro/impacta-alpha-despensa/pkg/uow"
)

type UpdateProductInput struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type UpdateProductOutput struct {
	RemainingQuantity int `json:"remaining_quantity"`
}

type ProductBalanceUpdateUseCase struct {
	Uow uow.UowInterface
}

var (
	ErrInternal = errors.New("internal error")
)

func NewProductBalanceUpdateUseCase(uow uow.UowInterface) *ProductBalanceUpdateUseCase {
	return &ProductBalanceUpdateUseCase{
		Uow: uow,
	}
}

func (p *ProductBalanceUpdateUseCase) Execute(input UpdateProductInput) (*UpdateProductOutput, error) {
	output := &UpdateProductOutput{}
	ctx := context.Background()

	err := p.Uow.Do(ctx, func(_ *uow.Uow) error {
		stockGateway, err := p.getStockRepository(ctx)
		if err != nil {
			log.Printf("stock gateway instantiate error: %s", err.Error())
			return ErrInternal
		}
		productBalanceGateway, err := p.getProductBalanceRepository(ctx)
		if err != nil {
			log.Printf("product balance gateway instantiate error: %s", err.Error())
			return ErrInternal
		}

		product, err := stockGateway.GetByID(input.ProductID)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				return errors.New("product not found")
			}
			return err
		}
		if product.Quantity == 0 {
			return errors.New("product not found")
		}

		err = product.UpdateQuantity(input.Quantity)
		if err != nil {
			if err == entity.ErrInsufficientStock {
				return err
			}
			log.Printf("update quantity error: %s", err.Error())
			return ErrInternal
		}
		err = stockGateway.UpdateQuantity(product)
		if err != nil {
			log.Printf("update quantity database error: %s", err.Error())
			return ErrInternal
		}

		productBalance := entity.NewProductBalance(input.ProductID, input.Quantity, product.Quantity)
		err = productBalanceGateway.Save(productBalance)
		if err != nil {
			log.Printf("product save error: %s", err.Error())
			return ErrInternal
		}

		output.RemainingQuantity = product.Quantity
		return nil
	})

	if err != nil {
		//in case of any error but product not found try to get the remaining quantity to send it to the user
		if err.Error() == "product not found" {
			return nil, err
		}
		currentCtx := context.Background()
		errUow := p.Uow.Do(ctx, func(_ *uow.Uow) error {
			productGateway, trErr := p.getStockRepository(currentCtx)
			if trErr != nil {
				log.Printf("error retrieving product quantity: %s", trErr.Error())
				return err
			}
			prd, idErr := productGateway.GetByID(input.ProductID)
			if idErr != nil {
				log.Printf("error retrieving product quantity: %s", idErr.Error())
				return err
			}
			output.RemainingQuantity = prd.Quantity
			return nil
		})
		if errUow != nil {
			log.Printf("error in unit of work %s", errUow.Error())
		}
		return output, err
	}
	return output, nil
}

func (p *ProductBalanceUpdateUseCase) getStockRepository(ctx context.Context) (gateway.StockGateway, error) {
	repo, err := p.Uow.GetRepository(ctx, "StockDb")
	if err != nil {
		return nil, err
	}
	return repo.(gateway.StockGateway), nil
}

func (p *ProductBalanceUpdateUseCase) getProductBalanceRepository(ctx context.Context) (gateway.ProductBalanceGateway, error) {
	repo, err := p.Uow.GetRepository(ctx, "ProductBalanceDB")
	if err != nil {
		return nil, err
	}
	return repo.(gateway.ProductBalanceGateway), nil
}
