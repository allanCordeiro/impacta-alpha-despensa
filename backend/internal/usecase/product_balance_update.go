package usecase

import (
	"context"

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
			return err
		}
		productBalanceGateway, err := p.getProductBalanceRepository(ctx)
		if err != nil {
			return err
		}

		product, err := stockGateway.GetByID(input.ProductID)
		if err != nil {
			return err
		}

		err = product.UpdateQuantity(input.Quantity)
		if err != nil {
			return err
		}
		err = stockGateway.UpdateQuantity(product)
		if err != nil {
			return err
		}

		productBalance := entity.NewProductBalance(input.ProductID, input.Quantity)
		err = productBalanceGateway.Save(productBalance)
		if err != nil {
			return err
		}

		output.RemainingQuantity = product.Quantity
		return nil
	})

	if err != nil {
		ctx := context.Background()
		stockGateway, errGetQty := p.getStockRepository(ctx)
		if errGetQty != nil {
			return nil, errGetQty
		}

		product, errGetQty := stockGateway.GetByID(input.ProductID)
		if errGetQty != nil {
			return nil, errGetQty
		}

		output.RemainingQuantity = product.Quantity
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
