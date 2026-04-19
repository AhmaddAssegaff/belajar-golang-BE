package server

import (
	"context"
	"database/sql"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"

	"belajar-go-be/internal/product"
	productdb "belajar-go-be/internal/product/sqlc"

	_ "belajar-go-be/docs"
)

type App struct {
	DB *sql.DB
	E  *echo.Echo
}

func New(db *sql.DB) *App {
	e := echo.New()

	queries := productdb.New(db)

	repo := product.NewRepository(queries)
	service := product.NewService(repo)
	handler := product.NewHandler(service)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "ok")
	})

	e.GET("/products", handler.GetProducts)

	return &App{
		DB: db,
		E:  e,
	}
}

func (a *App) Start(addr string) error {
	return a.E.Start(addr)
}

func (a *App) Stop(ctx context.Context) error {
	return a.E.Shutdown(ctx)
}
