package database

import (
	"database/sql"

	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
)

type StockDb struct {
	DB *sql.DB
}

func NewStockDb(db *sql.DB) *StockDb {
	return &StockDb{
		DB: db,
	}
}

func (s *StockDb) Save(product *entity.Product) error {
	stmt, err := s.DB.Prepare(
		"INSERT INTO stock_products(id, name, creation_date, quantity, expiration_date) VALUES ($1, $2, $3, $4, $5)",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.ID, product.Name, product.CreationDate, product.Quantity, product.ExpirationDate)
	if err != nil {
		return err
	}
	return nil
}

func (s *StockDb) GetByID(id string) (*entity.Product, error) {
	product := &entity.Product{}
	stmt, err := s.DB.Prepare("SELECT id, name, creation_date, quantity, expiration_date FROM stock_products WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&product.ID,
		&product.Name,
		&product.CreationDate,
		&product.Quantity,
		&product.ExpirationDate,
	)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *StockDb) GetAllProducts() ([]entity.Product, error) {
	rows, err := s.DB.Query("SELECT id, name, creation_date, quantity, expiration_date FROM stock_products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product
	for rows.Next() {
		var prod entity.Product
		err = rows.Scan(&prod.ID, &prod.Name, &prod.CreationDate, &prod.Quantity, &prod.ExpirationDate)
		if err != nil {
			return nil, err
		}
		products = append(products, prod)
	}
	return products, nil
}

func (s *StockDb) UpdateQuantity(stock *entity.Product) error {
	stmt, err := s.DB.Prepare("UPDATE stock_products SET quantity = $2 WHERE id = $1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(stock.ID, stock.Quantity)
	if err != nil {
		return err
	}
	return nil
}
