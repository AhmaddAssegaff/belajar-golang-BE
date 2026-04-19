package main

import (
	"log/slog"

	"belajar-go-be/config"

	_ "belajar-go-be/docs"

	echoSwagger "github.com/swaggo/echo-swagger"

	"belajar-go-be/internal/product"
	productdb "belajar-go-be/internal/product/sqlc"
	"database/sql"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	db, err := sql.Open("postgres", cfg.Database.DbUrl)
	if err != nil {
		slog.Error("db error", "err", err)
		return
	}

	if err := db.Ping(); err != nil {
		slog.Error("db ping error", "err", err)
		return
	}

	queries := productdb.New(db)

	repo := product.NewRepository(queries)
	service := product.NewService(repo)
	handler := product.NewHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/products", handler.GetProducts)

	if err := e.Start(cfg.AppHost); err != nil {
		slog.Error("failed to start server", "error", err)
	}
}
