package main

import (
	"log/slog"
	"net/http"

	"belajar-go-be/config"
	"database/sql"

	productHttp "belajar-go-be/internal/product/adapter/http"
	productRepo "belajar-go-be/internal/product/adapter/repository"
	productApp "belajar-go-be/internal/product/application"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
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

	// Echo instance
	e := echo.New()

	repo := productRepo.NewProductRepository(db)
	service := productApp.NewProductService(repo)
	handler := productHttp.NewProductHandler(service)

	// Middleware
	e.Use(middleware.RequestLogger()) // use the RequestLogger middleware with slog logger
	e.Use(middleware.Recover())       // recover panics as errors for proper error handling

	// Routes
	e.GET("/", hello)
	e.GET("/products", handler.GetAll)

	// Start server
	if err := e.Start(cfg.AppHost); err != nil {
		slog.Error("failed to start server", "error", err)
	}
}

// Handler
func hello(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
