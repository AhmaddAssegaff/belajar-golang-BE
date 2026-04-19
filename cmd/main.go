package main

import (
	"database/sql"
	"log/slog"

	"belajar-go-be/config"
	"belajar-go-be/internal/server"

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

	app := server.New(db)

	if err := app.Start(cfg.AppHost); err != nil {
		slog.Error("failed to start server", "error", err)
	}
}
