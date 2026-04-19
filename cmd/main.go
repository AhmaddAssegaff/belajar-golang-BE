package main

import (
	"log/slog"

	"belajar-go-be/config"
	"database/sql"

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

	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	if err := e.Start(cfg.AppHost); err != nil {
		slog.Error("failed to start server", "error", err)
	}
}
