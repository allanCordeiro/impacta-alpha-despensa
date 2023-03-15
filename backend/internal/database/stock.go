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

func (s *StockDb) Save(product entity.Product) error {
	stmt, err := s.DB.Prepare(
		"INSERT INTO stock(id, name, creation_date, quantity, expiration_date) VALUES (?, ?, ?, ?, ?)",
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
	stmt, err := s.DB.Prepare("SELECT id, name, creation_date, quantity, expiration_date FROM stock WHERE id = ?")
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
