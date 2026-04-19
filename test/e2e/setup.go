package e2e

import (
	"belajar-go-be/internal/server"
	"context"
	"database/sql"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	cfg "belajar-go-be/test/e2e/config"
)

func waitForServer(url string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	delay := 100 * time.Millisecond

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		resp, err := http.Get(url)
		if err == nil && resp.StatusCode == 200 {
			resp.Body.Close()
			return nil
		}

		time.Sleep(delay)

		if delay < time.Second {
			delay *= 2
		}
	}
}

func StartTestServer() (*server.App, *sql.DB, func()) {
	db, err := sql.Open("postgres", cfg.TestDB)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	app := server.New(db)

	go func() {
		_ = app.Start(":8081")
	}()

	_ = waitForServer(cfg.HealthURL)

	cleanup := func() {
		db.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_ = app.Stop(ctx)
	}

	return app, db, cleanup
}
