package mock_db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func SetupDB() *sql.DB {
	db, err := sql.Open("sqlite3", "file:foobar?mode=memory&cache=shared")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE stock_products (id varchar(255), name varchar(255), creation_date date, quantity int, expiration_date date, created_at timestamp DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE product_balance (id int PRIMARY KEY, product_id varchar(255) NOT NULL, deducted_amount int NOT NULL, remaining_quantity int NOT NULL, deducted_date timestamp NOT NULL)")
	if err != nil {
		panic(err)
	}
	return db
}

func TearDown(db *sql.DB) {
	defer db.Close()

	db.Exec("DROP TABLE stock_products")
	db.Exec("DROP TABLE product_balance")
}
