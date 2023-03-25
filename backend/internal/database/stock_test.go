package database

import (
	"database/sql"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/domain/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewStockDb(t *testing.T) {
	db := setupDB()
	clientDB := NewStockDb(db)
	t.Run("Given a valid product, when calls save method then data should be stored", func(t *testing.T) {
		expectedName := "product 1"
		expectedCreationDate := time.Now()
		expectedQuantity := 20
		expectedExpirationDate := time.Now().Add(time.Hour * 24 * 10)

		product := entity.NewProduct(expectedName, expectedCreationDate, expectedQuantity, expectedExpirationDate)
		prdOk, err := product.IsValid()
		assert.Nil(t, err)
		assert.True(t, prdOk)
		err = clientDB.Save(product)
		assert.Nil(t, err)

		receivedPrd, err := clientDB.GetByID(product.ID)
		assert.Nil(t, err)
		assert.Equal(t, expectedName, receivedPrd.Name)
		assert.True(t, expectedCreationDate.Equal(receivedPrd.CreationDate))
		assert.Equal(t, expectedQuantity, receivedPrd.Quantity)
		assert.True(t, expectedExpirationDate.Equal(receivedPrd.ExpirationDate))
	})
	tearDown(db)
}

func TestGetStock(t *testing.T) {
	db := setupDB()
	clientDB := NewStockDb(db)
	t.Run("Given a valid product, when calls save method then data should be stored", func(t *testing.T) {
		expectedPrd1Name := "product 1"
		expectedPrd1CreationDate := time.Now()
		expectedPrd1Quantity := 20
		expectedPrd1ExpirationDate := time.Now().Add(time.Hour * 24 * 10)
		product := entity.NewProduct(
			expectedPrd1Name,
			expectedPrd1CreationDate,
			expectedPrd1Quantity,
			expectedPrd1ExpirationDate)
		err := clientDB.Save(product)
		assert.Nil(t, err)
		expectedExpiredPrdName := "product 2"
		expectedExpiredPrdCreationDate := time.Now()
		expectedExpiredPrdQuantity := 20
		expectedExpiredPrdExpirationDate := time.Now().Add(time.Hour * 24 * 10)

		product = entity.NewProduct(
			expectedExpiredPrdName,
			expectedExpiredPrdCreationDate,
			expectedExpiredPrdQuantity,
			expectedExpiredPrdExpirationDate)
		err = clientDB.Save(product)
		assert.Nil(t, err)

		productList, err := clientDB.GetAllProducts()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(productList))
	})

	tearDown(db)
}

func setupDB() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE stock_products (id varchar(255), name varchar(255), creation_date date, quantity int, expiration_date date)")
	if err != nil {
		panic(err)
	}
	return db
}

func tearDown(db *sql.DB) {
	defer db.Close()

	db.Exec("DROP TABLE stock_products")
}
