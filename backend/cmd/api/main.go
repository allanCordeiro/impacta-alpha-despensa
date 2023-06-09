package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/AllanCordeiro/impacta-alpha-despensa/docs"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/database"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/webserver/handlers/balance_handlers"
	balancehandlers "github.com/AllanCordeiro/impacta-alpha-despensa/internal/webserver/handlers/balance_handlers"
	"github.com/AllanCordeiro/impacta-alpha-despensa/internal/webserver/handlers/stock_handlers"
	"github.com/AllanCordeiro/impacta-alpha-despensa/pkg/uow"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/subosito/gotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title 						Despensa Faculdade Impacta's Project
// @version 					1.0
// @description 				Product API with stock management
// @termsOfService 				http://www.swagger.io/terms

// @contact.name 				Allan Cordeiro
// @contact.url 				http://www.allancordeiro.com
// @contact.email 				allan.cordeiro.santos@gmail.com

// @host 						localhost:8000
// @basePath 					/
func main() {
	db, err := sql.Open("postgres", getEnvConfig("DB_URL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://sql/migrations",
		"postgres",
		driver)
	if err != nil {
		panic(err)
	}
	m.Up()

	stockDB := database.NewStockDb(db)
	stockHandler := stock_handlers.NewStockandler(stockDB)

	uow := uow.NewUow(context.Background(), db)
	uow.Register("StockDb", func(tx *sql.Tx) interface{} {
		return database.NewStockDb(db)
	})
	uow.Register("ProductBalanceDB", func(tx *sql.Tx) interface{} {
		return database.NewProductBalanceDB(db)
	})

	productBalancerHandler := balancehandlers.NewProductBalance(uow)

	productBalanceDB := database.NewProductBalanceDB(db)
	producBalancerHandlerWithGateway := balance_handlers.NewProductBalanceWithGateway(productBalanceDB, stockDB)

	docs.SwaggerInfo.Host = getEnvConfig("SWAGGER_HOST")
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	minQuantity, err := strconv.Atoi(getEnvConfig("MINIMAL_QUANTITY"))
	if err != nil {
		minQuantity = 0
	}

	r.Use(middleware.WithValue("min-quantity", minQuantity))
	//handlers
	r.Route("/api/stock", func(r chi.Router) {
		r.Post("/", stockHandler.CreateProduct)
		r.Get("/", stockHandler.GetProducts)
		r.Get("/statistics", stockHandler.GetStatistics)
	})

	r.Route("/api/products/balance", func(r chi.Router) {
		r.Get("/", producBalancerHandlerWithGateway.GetBalance)
	})
	r.Route("/api/products/{productID}", func(r chi.Router) {
		r.Put("/decrease", productBalancerHandler.CreateProductBalance)
		r.Get("/balance", producBalancerHandlerWithGateway.GetProductBalance)
	})

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getEnvConfig(config string) string {
	envVar := os.Getenv(config)
	if envVar == "" {
		err := gotenv.Load(".env")
		if err != nil {
			panic(".env file not found.")
		}
		envVar = os.Getenv(config)
	}
	if config == "" {
		panic("environment config not found")
	}
	return envVar
}
