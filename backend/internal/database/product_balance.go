package database

import (
	"database/sql"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
)

type ProductBalanceDB struct {
	DB *sql.DB
}

func NewProductBalanceDB(db *sql.DB) *ProductBalanceDB {
	return &ProductBalanceDB{
		DB: db,
	}
}

func (p *ProductBalanceDB) Save(productBalance *entity.ProductBalance) error {
	stmt, err := p.DB.Prepare("INSERT INTO product_balance(product_id, deducted_amount, deducted_date, remaining_quantity) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(productBalance.ProductID, productBalance.DeductedAmount, productBalance.DeductedDate, productBalance.RemainingQuantity)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductBalanceDB) GetByProductId(productId string) ([]entity.ProductBalance, error) {
	rows, err := p.DB.Query("SELECT product_id, deducted_amount, remaining_quantity, deducted_date FROM product_balance WHERE product_id = $1", productId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productsBalance []entity.ProductBalance
	for rows.Next() {
		var prod entity.ProductBalance
		err = rows.Scan(&prod.ProductID, &prod.DeductedAmount, &prod.RemainingQuantity, &prod.DeductedDate)
		if err != nil {
			return nil, err
		}
		productsBalance = append(productsBalance, prod)
	}
	return productsBalance, nil
}

func (p *ProductBalanceDB) GetAllProductsBalance() ([]entity.ProductBalance, error) {
	rows, err := p.DB.Query("SELECT product_id, deducted_amount, remaining_quantity, deducted_date FROM product_balance")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var productsBalance []entity.ProductBalance
	for rows.Next() {
		var prod entity.ProductBalance
		err = rows.Scan(&prod.ProductID, &prod.DeductedAmount, &prod.RemainingQuantity, &prod.DeductedDate)
		if err != nil {
			return nil, err
		}
		productsBalance = append(productsBalance, prod)
	}
	return productsBalance, nil
}
