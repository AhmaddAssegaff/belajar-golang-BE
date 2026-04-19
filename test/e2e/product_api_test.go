package e2e

import (
	"net/http"
	"testing"

	cfg "belajar-go-be/test/e2e/config"
)

func TestGetProducts(t *testing.T) {
	_, _, cleanup := StartTestServer()
	defer cleanup()

	resp, err := http.Get(cfg.BaseURL + "/products")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}
